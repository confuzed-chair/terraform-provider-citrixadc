resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}
resource "citrixadc_appfwprofile_jsondosurl_binding" "tf_binding" {
  name                        = citrixadc_appfwprofile.tf_appfwprofile.name
  jsondosurl                  = ".*"
  state                       = "ENABLED"
  alertonly                   = "ON"
  isautodeployed              = "AUTODEPLOYED"
  jsonmaxarraylengthcheck     = "ON"
  jsonmaxdocumentlengthcheck  = "ON"
  jsonmaxcontainerdepth       = 5
  jsonmaxobjectkeylengthcheck = "OFF"
  jsonmaxarraylength          = 100000
  jsonmaxdocumentlength       = 200000
  jsonmaxobjectkeycountcheck  = "ON"
  jsonmaxobjectkeylength      = 128
  jsonmaxobjectkeycount       = 1000
  jsonmaxstringlengthcheck    = "ON"
  jsonmaxcontainerdepthcheck  = "ON"
  jsonmaxstringlength         = 1000
  comment                     = "Testing"
}