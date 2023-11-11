package prepare

import (
	"fmt"

	clabernetesconstants "github.com/srl-labs/clabernetes/constants"
	clabernetesmanagertypes "github.com/srl-labs/clabernetes/manager/types"
	clabernetesmanagerutil "github.com/srl-labs/clabernetes/manager/util"
	clabernetesutil "github.com/srl-labs/clabernetes/util"

	claberneteserrors "github.com/srl-labs/clabernetes/errors"
	k8scorev1 "k8s.io/api/core/v1"
)

// certificates initializes clabernetes-manager certs -- this includes a certificate authority
// and client certs.
func certificates(c clabernetesmanagertypes.Clabernetes) error {
	clabernetesutil.MustCreateDirectory(
		clabernetesconstants.CertificateDirectory,
		clabernetesconstants.PermissionsEveryoneRead,
	)

	caDirectory := ensureCaDirectory()

	clientCertDirectory := ensureClientCertDirectory()

	webhookCertDirectory := ensureWebhookCertDirectory()

	certsExist := true

	for _, directory := range []string{caDirectory, clientCertDirectory} {
		for _, fileName := range []string{"ca.crt", "tls.crt", "tls.key"} {
			exists := clabernetesutil.MustFileExists(fmt.Sprintf("%s/%s", directory, fileName))
			if !exists {
				certsExist = false

				break
			}
		}

		if !certsExist {
			break
		}
	}

	if certsExist {
		return nil
	}

	secret, err := clabernetesmanagerutil.GetCertificatesSecret(c)
	if err != nil {
		return fmt.Errorf("getting certificates secret: %w", err)
	}

	if secret.Data == nil {
		return fmt.Errorf("%w: certificate data is nil", claberneteserrors.ErrPrepare)
	}

	caData, err := getCertFromSecret("ca", secret)
	if err != nil {
		return err
	}

	err = caData.Write(caDirectory)
	if err != nil {
		return fmt.Errorf("write ca certificate data: %w", err)
	}

	clientData, err := getCertFromSecret("client", secret)
	if err != nil {
		return err
	}

	err = clientData.Write(clientCertDirectory)
	if err != nil {
		return fmt.Errorf("write client certificate data: %w", err)
	}

	webhookData, err := getCertFromSecret("webhook", secret)
	if err != nil {
		return err
	}

	err = webhookData.Write(webhookCertDirectory)
	if err != nil {
		return fmt.Errorf("write webhook certificate data: %w", err)
	}

	return nil
}

func getCertFromSecret(
	certType string,
	secret *k8scorev1.Secret,
) (*clabernetesutil.CertData, error) {
	crt, crtOk := secret.Data[fmt.Sprintf("%s-ca.crt", certType)]
	tls, tlsOk := secret.Data[fmt.Sprintf("%s-tls.crt", certType)]
	key, keyOk := secret.Data[fmt.Sprintf("%s-tls.key", certType)]

	if !crtOk || !tlsOk || !keyOk {
		return nil, fmt.Errorf(
			"%w: missing one or more certificate data fields for cert type '%s'",
			claberneteserrors.ErrPrepare,
			certType,
		)
	}

	return &clabernetesutil.CertData{
		TLS: tls,
		CRT: crt,
		Key: key,
	}, nil
}

func ensureCaDirectory() string {
	caDirectory := fmt.Sprintf(
		"%s/%s",
		clabernetesconstants.CertificateDirectory,
		clabernetesconstants.CertificateAuthoritySubDir,
	)

	clabernetesutil.MustCreateDirectory(caDirectory, clabernetesconstants.PermissionsEveryoneRead)

	return caDirectory
}

func ensureClientCertDirectory() string {
	clientCertDirectory := fmt.Sprintf(
		"%s/%s",
		clabernetesconstants.CertificateDirectory,
		clabernetesconstants.ClientCertificateSubDir,
	)

	clabernetesutil.MustCreateDirectory(
		clientCertDirectory,
		clabernetesconstants.PermissionsEveryoneRead,
	)

	return clientCertDirectory
}

func ensureWebhookCertDirectory() string {
	webhookCertDirectory := fmt.Sprintf(
		"%s/%s",
		clabernetesconstants.CertificateDirectory,
		clabernetesconstants.WebhookCertificateSubDir,
	)

	clabernetesutil.MustCreateDirectory(
		webhookCertDirectory,
		clabernetesconstants.PermissionsEveryoneRead,
	)

	return webhookCertDirectory
}
