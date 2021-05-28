source = [
  "./dist/gitstat-macos_darwin_amd64/gitstat"]
bundle_id = "com.gitstat"

apple_id {
  username = "niels@kryger.nl"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Niels Krijger"
}