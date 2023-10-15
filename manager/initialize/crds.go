package initialize

import (
	"embed"
	"fmt"
	"path"

	clabernetesassets "github.com/srl-labs/clabernetes/assets"
	clabernetesmanagertypes "github.com/srl-labs/clabernetes/manager/types"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/equality"
	apimachineryerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func crds(c clabernetesmanagertypes.Clabernetes) error {
	loadedCrds, err := loadCrdsFromAssets()
	if err != nil {
		return err
	}

	extensionsClient, err := apiextensionsclient.NewForConfig(c.GetKubeConfig())
	if err != nil {
		return fmt.Errorf("creating apiextensions client: %w", err)
	}

	for _, crd := range loadedCrds {
		err = applyCRD(c, extensionsClient, crd)
		if err != nil {
			return fmt.Errorf("%s: %w", fmt.Sprintf("apply crd %s", crd.Name), err)
		}
	}

	return nil
}

func getCrdFilenamesFromAssets(fs embed.FS, dir string) ([]string, error) {
	objs, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, 0)

	for _, obj := range objs {
		if obj.IsDir() {
			continue
		}

		fileNames = append(fileNames, path.Join(dir, obj.Name()))
	}

	return fileNames, nil
}

func loadCrdsFromAssets() ([]*apiextensionsv1.CustomResourceDefinition, error) {
	crdFileNames, err := getCrdFilenamesFromAssets(clabernetesassets.Assets, "crd")
	if err != nil {
		return nil, fmt.Errorf("globbing crd assets: %w", err)
	}

	allCrds := make([]*apiextensionsv1.CustomResourceDefinition, len(crdFileNames))

	for i, crdFileName := range crdFileNames {
		b, readErr := clabernetesassets.Assets.ReadFile(crdFileName)
		if readErr != nil {
			return nil, fmt.Errorf("%s: %w", fmt.Sprintf("reading crd %s", crdFileName), readErr)
		}

		crd := &apiextensionsv1.CustomResourceDefinition{}

		// *note* apimachinery runtime yaml, dunno what's different but... whatever
		err = yaml.Unmarshal(b, crd)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fmt.Sprintf("unmarshalling crd %s", crdFileName), err)
		}

		allCrds[i] = crd
	}

	return allCrds, nil
}

func applyCRD(
	c clabernetesmanagertypes.Clabernetes,
	client *apiextensionsclient.Clientset,
	crd *apiextensionsv1.CustomResourceDefinition,
) error {
	ctx, ctxCancel := c.NewContextWithTimeout()
	defer ctxCancel()

	currentCrd, err := client.ApiextensionsV1().CustomResourceDefinitions().Get(
		ctx,
		crd.Name,
		metav1.GetOptions{},
	)

	if err != nil && apimachineryerrors.IsNotFound(err) {
		// crd didn't exist, create it
		_, err = client.ApiextensionsV1().CustomResourceDefinitions().Create(
			ctx,
			crd,
			metav1.CreateOptions{},
		)

		return err
	} else if err != nil {
		// something weird happened, bail out
		return err
	}

	if equality.Semantic.DeepEqual(currentCrd.Spec.Versions, crd.Spec.Versions) &&
		equality.Semantic.DeepEqual(currentCrd.Spec.Conversion, crd.Spec.Conversion) {
		// crds are the same, nothing to do
		return nil
	}

	// update the versions
	currentCrd.Spec.Versions = crd.Spec.Versions

	// update the conversion bits too
	currentCrd.Spec.Conversion = crd.Spec.Conversion

	_, err = client.ApiextensionsV1().CustomResourceDefinitions().Update(
		ctx,
		currentCrd,
		metav1.UpdateOptions{},
	)

	return err
}
