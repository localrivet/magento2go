# Magento2Go

Connects go to a Magento Community REST Api using swagger to generate the the go client and structs.

## Issue With Magento's Api

Magento's REST Api has on major issue:

*It doesn't actually follow it's own swagger documentation and requires `fixing` the raw json response string before marshalling it into a corresponding struct.* 

## Testing

All tests are written using `ginkgo` 
- https://github.com/onsi/ginkgo

Once your instance is up and running and you're ready to run the test run the following command.

```sh
ginkgo watch
# or
go test
```

Due to Magento's Api issues mocking the Api using their own specs for testing is not recommended as it will not match the actual output of Magento's Api. 

Testing requires a locally or remotely run Magento dev instance. These tests are ran again a locally run Magento docker instance.

Here's how I am setup:


1. I recommend using [Mark Shust's Docker Configuration for Magento](https://github.com/markshust/docker-magento)
He's done a fantastic job of putting this together and it makes running a local dev environment a breeze.

    Setup is easy:
    ```sh
    curl -s https://raw.githubusercontent.com/markshust/docker-magento/master/lib/onelinesetup | bash -s -- magento.test 2.4.3-p1
    ```

2. You're going to need a `username` and `password` to get it setup properly. They can be found at Adobe's Magento site.

    - Login or create an Adobe account.
    https://account.magento.com/customer/account/login/

    - Once logged in navigate to this URL and click `Create A New Access Key`.
    https://marketplace.magento.com/customer/accessKeys/


3. Setup a the local testing URL `magento.test` in your `/ect/hosts` file.
    ```sh
    127.0.0.1   magento.test
    ```

4. Install your Magento's cron. This will save you a lot of headaches and battles with Magento's cache.
   ```sh
   bin/magento cron:install 
   ```
   

5. Login to your local Magento instance and create your access token

    Login > System > Extensions > Integrations > Add New Integration

    On a local dev instance you can safely give permission for everything. *Don't do this in production*


6. Create a `.env` file for testing.

    ```sh
    MAGENTO_STORE_SCHEME=https
    MAGENTO_STORE_HOSTNAME=magento.test
    MAGENTO_STORE_CODE=default
    MAGENTO_ACCESS_TOKEN=<put your Magentos access token here>
    ```


### Regenerating Swagger Client

Generate the Api client using swagger
```sh
swagger generate client -f magento.schema.json -a magento2go --default-consumes application/json
```

## Usage

```sh
var err error
var client *client.MagentoCommunity

// create the client
client, err := NewClient(&struct {
    accessToken string
    host        string
    path        string
    scheme      string
    debug       bool
}{
    accessToken: os.Getenv("MAGENTO_ACCESS_TOKEN"),
    host:        os.Getenv("MAGENTO_STORE_HOSTNAME"),
    path:        "/rest/default",
    scheme:      os.Getenv("MAGENTO_STORE_SCHEME"),
    debug:       false,
})

if err != nil {
    fmt.Print(err)
}

// create the api
// timeout in seconds
var timeout int64 = 300 
api = NewMagentoApi(client, timeout)

// get product by sku
var product *models.CatalogDataProductInterface
product, err = api.Product.GetProductBySku("valid-sku")
if err != nil {
    fmt.Print(err)
}

// do something with product
// Attribute set id
product.AttributeSetID // int64 

// Created date
product.CreatedAt // string 

// Custom attributes values.
product.CustomAttributes // []*FrameworkAttributeInterface 

// extension attributes
product.ExtensionAttributes // *CatalogDataProductExtensionInterface 

// Id
product.ID // int64

// Media gallery entries
product.MediaGalleryEntries // []*CatalogDataProductAttributeMediaGalleryEntryInterface 

// Name
product.Name // string

// List of product options
product.Options // []*CatalogDataProductCustomOptionInterface 

// Price
product.Price // float64 

// Product links info
product.ProductLinks // []*CatalogDataProductLinkInterface

// Sku
// Required: true
product.Sku // *string 

// Status
product.Status // int64 

// List of product tier prices
product.TierPrices // []*CatalogDataProductTierPriceInterface 

// Type id
product.TypeID // string

// Updated date
product.UpdatedAt // string

// Visibility
product.Visibility // int64 

// Weight
product.Weight // float64
```

## Api Methods Planned

### Products
- [x] GetProductBySku
- [x] GetAllProducts


### Orders & more
- Coming soon

