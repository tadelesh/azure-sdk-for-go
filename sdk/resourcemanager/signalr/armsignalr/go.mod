module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr

go 1.18

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.22.1-0.20220315231014-ed309e73db6b
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.13.2
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.9.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v0.3.1
)

require (
	github.com/AzureAD/microsoft-authentication-library-for-go v0.4.0 // indirect
	github.com/dnaeon/go-vcr v1.1.0 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20210115035449-ce105d075bb4 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/azcore => ../../../../../../Azure/azure-sdk-for-go/sdk/azcore
