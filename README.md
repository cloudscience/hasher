# Hasher
A go-crypto based command line utility for generating cryptographic hashes of files.

```
hasher <file1> <file2> <file3> <filen>
```

hasher will print formatted JSON to STDOUT, e.g.:

```
{
  "SHA256hash": "769f1bc82e93f7dd6241c8b1c2e89c939d18eaec256be12513c38bc2bfde2489",
  "SHA512hash": "e8b80858d927031fd5c5e892d260f44e15e4abe868c8f0aa11763ce64f23781adc2665842a9a68552b7b0e7cddce4d3a8df0e83455c61c780e42d1e83690b1d6",
  "MD5hash": "bb065a9369f76dceb78e0f3df19c7551",
  "Filename": "file-test.json",
  "Filepath": "/home/username/file/"
}
```
