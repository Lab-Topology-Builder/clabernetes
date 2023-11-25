package reconciler_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"

	clabernetesapistopology "github.com/srl-labs/clabernetes/apis/topology"
	clabernetesapistopologyv1alpha1 "github.com/srl-labs/clabernetes/apis/topology/v1alpha1"
	clabernetesconfig "github.com/srl-labs/clabernetes/config"
	clabernetesconstants "github.com/srl-labs/clabernetes/constants"
	clabernetescontrollerstopologyreconciler "github.com/srl-labs/clabernetes/controllers/topology/reconciler"
	claberneteslogging "github.com/srl-labs/clabernetes/logging"
	clabernetestesthelper "github.com/srl-labs/clabernetes/testhelper"
	clabernetesutil "github.com/srl-labs/clabernetes/util"
	clabernetesutilcontainerlab "github.com/srl-labs/clabernetes/util/containerlab"
	k8scorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const renderPersistentVolumeClaimTestName = "persistentvolumeclaim/render-pvc"

func TestResolvePersistentVolumeClaim(t *testing.T) {
	cases := []struct {
		name                 string
		ownedPVCs            *k8scorev1.PersistentVolumeClaimList
		clabernetesConfigs   map[string]*clabernetesutilcontainerlab.Config
		owningTopologyObject clabernetesapistopologyv1alpha1.TopologyCommonObject
		expectedCurrent      []string
		expectedMissing      []string
		expectedExtra        []*k8scorev1.PersistentVolumeClaim
	}{
		{
			name:               "simple",
			ownedPVCs:          &k8scorev1.PersistentVolumeClaimList{},
			clabernetesConfigs: nil,
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled: true,
						},
					},
				},
			},
			expectedCurrent: nil,
			expectedMissing: nil,
			expectedExtra:   []*k8scorev1.PersistentVolumeClaim{},
		},
		{
			name: "extra-pvcs",
			ownedPVCs: &k8scorev1.PersistentVolumeClaimList{
				Items: []k8scorev1.PersistentVolumeClaim{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "extra-pvc-1",
							Namespace: "clabernetes",
							Labels: map[string]string{
								clabernetesconstants.LabelTopologyNode: "node2",
							},
						},
					},
				},
			},
			clabernetesConfigs: nil,
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled: true,
						},
					},
				},
			}, expectedCurrent: nil,
			expectedMissing: nil,
			expectedExtra: []*k8scorev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "extra-pvc-1",
						Namespace: "clabernetes",
						Labels: map[string]string{
							clabernetesconstants.LabelTopologyNode: "node2",
						},
					},
				},
			},
		},
		{
			name:      "missing-pvcs",
			ownedPVCs: &k8scorev1.PersistentVolumeClaimList{},
			clabernetesConfigs: map[string]*clabernetesutilcontainerlab.Config{
				"node1": nil,
				"node2": nil,
			},
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled: true,
						},
					},
				},
			},
			expectedCurrent: nil,
			expectedMissing: []string{"node1", "node2"},
			expectedExtra:   []*k8scorev1.PersistentVolumeClaim{},
		},
		{
			name: "pvcs-but-persistence-disabled",
			ownedPVCs: &k8scorev1.PersistentVolumeClaimList{
				Items: []k8scorev1.PersistentVolumeClaim{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "extra-pvc-1",
							Namespace: "clabernetes",
							Labels: map[string]string{
								clabernetesconstants.LabelTopologyNode: "node2",
							},
						},
					},
				},
			},
			clabernetesConfigs: map[string]*clabernetesutilcontainerlab.Config{
				"node1": nil,
				"node2": nil,
			},
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled: false,
						},
					},
				},
			},
			expectedCurrent: nil,
			expectedMissing: nil,
			expectedExtra: []*k8scorev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "extra-pvc-1",
						Namespace: "clabernetes",
						Labels: map[string]string{
							clabernetesconstants.LabelTopologyNode: "node2",
						},
					},
				},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Logf("%s: starting", testCase.name)

				reconciler := clabernetescontrollerstopologyreconciler.NewPersistentVolumeClaimReconciler(
					&claberneteslogging.FakeInstance{},
					clabernetesapistopology.Containerlab,
					clabernetesconfig.GetFakeManager,
				)

				got, err := reconciler.Resolve(
					testCase.ownedPVCs,
					testCase.clabernetesConfigs,
					testCase.owningTopologyObject,
				)
				if err != nil {
					t.Fatal(err)
				}

				var gotCurrent []string

				for current := range got.Current {
					gotCurrent = append(gotCurrent, current)
				}

				if !clabernetesutil.StringSliceContainsAll(gotCurrent, testCase.expectedCurrent) {
					clabernetestesthelper.FailOutput(t, gotCurrent, testCase.expectedCurrent)
				}

				if !clabernetesutil.StringSliceContainsAll(got.Missing, testCase.expectedMissing) {
					clabernetestesthelper.FailOutput(t, got.Missing, testCase.expectedMissing)
				}

				clabernetestesthelper.MarshaledEqual(t, got.Extra, testCase.expectedExtra)
			})
	}
}

func TestRenderPersistentVolumeClaim(t *testing.T) {
	cases := []struct {
		name                 string
		owningTopologyObject clabernetesapistopologyv1alpha1.TopologyCommonObject
		clabernetesConfigs   map[string]*clabernetesutilcontainerlab.Config
		nodeName             string
	}{
		{
			name: "simple",
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pvc-test",
					Namespace: "clabernetes",
				},
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled: true,
						},
					},
				},
			},
			clabernetesConfigs: map[string]*clabernetesutilcontainerlab.Config{
				"node1": nil,
			},
			nodeName: "node1",
		},
		{
			name: "explicit-storage-class",
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pvc-test",
					Namespace: "clabernetes",
				},
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled:          true,
							StorageClassName: "my-custom-storage-class",
						},
					},
				},
			},
			clabernetesConfigs: map[string]*clabernetesutilcontainerlab.Config{
				"node1": nil,
			},
			nodeName: "node1",
		},
		{
			name: "explicit-claim-size",
			owningTopologyObject: &clabernetesapistopologyv1alpha1.Containerlab{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pvc-test",
					Namespace: "clabernetes",
				},
				Spec: clabernetesapistopologyv1alpha1.ContainerlabSpec{
					TopologyCommonSpec: clabernetesapistopologyv1alpha1.TopologyCommonSpec{
						Persistence: clabernetesapistopologyv1alpha1.Persistence{
							Enabled:   true,
							ClaimSize: "99Gi",
						},
					},
				},
			},
			clabernetesConfigs: map[string]*clabernetesutilcontainerlab.Config{
				"node1": nil,
			},
			nodeName: "node1",
		},
	}

	for _, testCase := range cases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Logf("%s: starting", testCase.name)

				reconciler := clabernetescontrollerstopologyreconciler.NewPersistentVolumeClaimReconciler(
					&claberneteslogging.FakeInstance{},
					clabernetesapistopology.Containerlab,
					clabernetesconfig.GetFakeManager,
				)

				got := reconciler.Render(
					testCase.owningTopologyObject,
					testCase.nodeName,
				)

				if *clabernetestesthelper.Update {
					clabernetestesthelper.WriteTestFixtureJSON(
						t,
						fmt.Sprintf(
							"golden/%s/%s.json",
							renderPersistentVolumeClaimTestName,
							testCase.name,
						),
						got,
					)
				}

				var want k8scorev1.PersistentVolumeClaim

				err := json.Unmarshal(
					clabernetestesthelper.ReadTestFixtureFile(
						t,
						fmt.Sprintf(
							"golden/%s/%s.json",
							renderPersistentVolumeClaimTestName,
							testCase.name,
						),
					),
					&want,
				)
				if err != nil {
					t.Fatal(err)
				}

				clabernetestesthelper.MarshaledEqual(t, got, want)
			})
	}
}

func TestPersistentVolumeClaimConforms(t *testing.T) {
	cases := []struct {
		name     string
		existing *k8scorev1.PersistentVolumeClaim
		rendered *k8scorev1.PersistentVolumeClaim
		ownerUID apimachinerytypes.UID
		conforms bool
	}{
		{
			name: "simple",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: true,
		},

		// claim size

		{
			name: "claim-size-conforms",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
				Spec: k8scorev1.PersistentVolumeClaimSpec{
					Resources: k8scorev1.ResourceRequirements{
						Requests: k8scorev1.ResourceList{
							"storage": resource.MustParse("1Gi"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				Spec: k8scorev1.PersistentVolumeClaimSpec{
					Resources: k8scorev1.ResourceRequirements{
						Requests: k8scorev1.ResourceList{
							"storage": resource.MustParse("1Gi"),
						},
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: true,
		},
		{
			name: "claim-size-mismatch",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
				Spec: k8scorev1.PersistentVolumeClaimSpec{
					Resources: k8scorev1.ResourceRequirements{
						Requests: k8scorev1.ResourceList{
							"storage": resource.MustParse("1Gi"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				Spec: k8scorev1.PersistentVolumeClaimSpec{
					Resources: k8scorev1.ResourceRequirements{
						Requests: k8scorev1.ResourceList{
							"storage": resource.MustParse("99Gi"),
						},
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: false,
		},

		// object meta annotations

		{
			name: "missing-clabernetes-global-annotations",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"somethingelse": "xyz",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"user-provided-global-annotation": "expected-value",
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: false,
		},
		{
			name: "clabernetes-global-annotations-wrong-value",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"user-provided-global-annotation": "xyz",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"user-provided-global-annotation": "expected-value",
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: false,
		},
		{
			name: "extra-annotations-ok",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"somethingelseentirely": "thisisok",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: true,
		},

		// object meta labels

		{
			name: "missing-clabernetes-global-annotations",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"somethingelse": "xyz",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"user-provided-global-label": "expected-value",
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: false,
		},
		{
			name: "clabernetes-global-labels-wrong-value",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"user-provided-global-label": "xyz",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"user-provided-global-label": "expected-value",
					},
				},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: false,
		},
		{
			name: "extra-labels-ok",
			existing: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"somethingelseentirely": "thisisok",
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							UID: apimachinerytypes.UID("clabernetes-testing"),
						},
					},
				},
			},
			rendered: &k8scorev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{},
			},
			ownerUID: apimachinerytypes.UID("clabernetes-testing"),
			conforms: true,
		},
	}

	for _, testCase := range cases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Logf("%s: starting", testCase.name)

				reconciler := clabernetescontrollerstopologyreconciler.NewPersistentVolumeClaimReconciler(
					&claberneteslogging.FakeInstance{},
					clabernetesapistopology.Containerlab,
					clabernetesconfig.GetFakeManager,
				)

				actual := reconciler.Conforms(
					testCase.existing,
					testCase.rendered,
					testCase.ownerUID,
				)
				if actual != testCase.conforms {
					clabernetestesthelper.FailOutput(
						t,
						testCase.existing,
						testCase.rendered,
					)
				}
			})
	}
}
