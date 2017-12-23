package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".DS_Store": []string{  // all .DS_Store assets.
        
          ".DS_Store",
        
      },
    
      ".tml": []string{  // all .tml assets.
        
          "jwt-api-http.tml",
        
          "jwt-api-readme.tml",
        
          "jwt-api.tml",
        
          "jwt-mock.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        ".DS_Store": { // all .DS_Store assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00"),
          path: ".DS_Store",
          root: ".DS_Store",
        },
      
    
      
        "jwt-api-http.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x73\xe3\x36\x12\x3e\x8b\xbf\xa2\xc3\x43\x42\x66\x35\x54\xed\x95\x29\x1f\xfc\x50\xbc\x9e\x49\x64\x97\xed\xd9\xa9\x3d\x6d\xc1\x64\x4b\x82\x4d\x11\x0c\x00\x5a\xa3\xb8\xfc\xdf\xb7\x1a\x00\x1f\xd2\x48\xb2\x64\x5b\xb2\x27\x1b\x1f\x46\x12\xd1\xe8\x77\x7f\x8d\x07\xe7\x9e\x49\x08\xbc\x4e\xaf\x07\xc7\x22\xd7\xf8\x55\x7f\xc2\xd9\xc7\x2f\xd7\x87\xa5\x1e\x1f\x67\x8c\x4f\x20\xc5\x21\xcf\x51\x41\x62\x87\xa1\x54\x98\x82\x16\xa0\xb4\x90\x08\xb7\x53\x0d\x09\xd1\x19\x16\x12\x13\x71\x8f\x12\x53\x18\x4a\x31\x01\x06\xac\xd4\x63\xcc\x35\x4f\x98\xe6\x22\x07\x89\x7f\x94\xa8\x74\xe4\x75\x56\x08\x6b\xfd\x1d\x54\x12\x3f\xe1\x2c\xf0\x89\xd1\x87\xdb\xa9\xfe\x60\x84\xf9\xa1\x17\x7a\x5e\xaf\xd7\x22\x31\x5f\x19\xcf\x15\x14\x52\xdc\xf3\x14\x53\xb8\xc3\x19\xdc\xb3\xac\x44\x18\x0a\x09\x89\x44\xa6\x79\x3e\xaa\x0d\xb9\x61\xca\x11\xe5\x6c\x82\x2a\xf2\xf4\xac\xc0\x36\x47\xa5\x25\xcf\x47\x46\xce\x95\xf9\x0a\x12\x75\x29\x73\xe5\x46\x2c\x6d\xae\xa1\x90\x3c\xd7\xec\x26\x73\x82\x2c\x87\x3b\x9c\x45\xde\xb0\xcc\x13\x08\x92\x16\xd7\xd0\xf1\x0a\xc2\x8a\xcb\x83\xd7\xb1\x7c\xc1\xff\xf8\xe5\xfa\x13\xd7\xb5\x86\x77\x38\x8b\xc1\x87\x7f\x38\xca\x20\x09\xbd\x47\xa3\x0e\xf9\xec\xe2\xac\x0e\x0d\x23\x82\x32\xd1\x30\x1d\xf3\x64\x0c\x63\x91\xa5\x0a\xf4\x18\x61\xac\x75\x01\xac\xe0\x30\x66\x79\x9a\xa1\x54\x46\x41\xeb\x20\x12\x7d\x7c\xf9\xf9\x84\xf8\x89\x02\xa5\x89\x90\x25\xa0\xa9\xb5\x17\x1f\x1e\xfe\x28\x85\x46\x88\xae\x8c\x8c\xe8\xfc\xe6\x16\x13\x1d\x0d\xd8\x04\xcd\x3f\x8f\x8f\x40\x8e\x73\xee\x73\x9a\x59\x7d\x1e\xbc\xce\x04\xb5\xe4\x89\x02\xf7\x19\xfd\x6e\x3f\xbd\x8e\x15\x29\x24\x9c\xa5\x94\x20\x7a\x76\xc4\x92\x3b\xcc\x53\x67\xe1\x00\xa7\xb5\xb7\x19\xe4\x38\xad\x38\xf3\x5c\x69\x96\x27\x08\xa5\x22\x0b\xe6\x54\xad\x79\xb2\x3c\x25\x26\x56\xa8\x8b\xc2\x00\xa7\xc1\x64\x51\x8f\x2e\xdc\x58\xb1\x8b\x6a\x84\xf0\xb3\x13\xd8\x84\xe7\x47\xfb\xe4\xc1\xeb\x54\x66\xc5\x30\xe9\x7a\x9d\xda\x96\xb8\xe2\xd6\xf5\x3a\x8f\xce\x90\xb3\x7c\x28\xa8\x2e\x90\xdf\x53\xa4\x72\x1b\x13\x57\x08\x54\x47\x23\xd4\xa6\x6e\x64\x0a\x9c\x68\x29\x00\x2c\xcb\x80\xdd\x33\x9e\x99\xa4\xb2\xa3\x0a\xc4\xd0\x78\x7a\x83\x88\x44\x5e\xaf\x47\xc2\x2f\x45\xa9\x31\x86\xde\x83\xf9\xf2\xd8\x23\x01\xf4\xfc\x77\xd4\x63\x91\xc6\x70\x36\xf8\xf5\xdc\xd0\xf5\xaf\x2e\xce\x07\x57\xfd\x0f\x47\xe7\x27\xff\x89\xe1\xe3\xd5\xf9\xc0\xa5\x2e\x25\x8f\xf3\x44\x68\x6c\x09\x12\xfd\x15\x7e\x26\x23\x4a\xcd\xb3\xc8\xd5\x71\x08\x28\xa5\x90\xe4\xac\x09\xc4\x07\xb5\x9b\x07\x38\xbd\x96\x2c\xc1\x20\xf4\x3a\x29\x0e\x51\x52\x36\x46\xd5\x68\x7f\xc2\x75\x50\xfd\x30\xcc\x7d\x2b\xca\xfc\xf0\xc3\x6e\xcd\xe7\x0b\xd7\x63\xcb\x68\x12\xf5\xf3\x34\x08\xc3\xd0\xf3\x3a\x89\xfe\x1a\xfd\x0b\x59\x8a\x32\x08\xa3\x2b\xd4\x81\x7f\x6c\x2b\xf2\xc3\xf5\xac\x40\xbf\x0b\x3e\x2b\x8a\xcc\x81\x4f\xef\x56\x89\xdc\xa7\x69\x4f\xa8\xe0\x22\x66\x03\xe4\x22\x97\xfa\x61\xb7\xad\xcb\xaf\x1c\xb3\x54\xd5\x13\xcd\x4f\xca\x0b\xbf\x94\x99\x1f\x03\x29\x76\x69\x39\x04\x61\xf4\xf9\xf2\xb7\xa8\x2a\x7a\xca\x0c\xd2\xdd\xeb\x68\xa1\x59\xd6\x25\xbf\x91\xc3\x48\xa7\x2a\x8d\xa2\x63\x51\xe6\x9a\x1c\x5d\xb9\x37\x08\x43\xaf\xc3\x87\x86\xf8\x87\x03\xc8\x79\x46\x9e\x5e\x6d\x48\x9f\x82\x31\x0c\xfc\x5f\x19\xcf\x2c\x5a\x53\x96\x55\x29\x7e\x69\xb3\xcd\x25\x5d\x42\xc2\x36\x34\xaf\xe3\x9b\x30\xfb\x31\x69\xd2\x35\x0f\x36\x30\xd8\x5a\xdc\xa9\xca\x08\xa5\xa4\xf2\xa8\x0d\x8a\x0f\xcc\x74\x4a\xba\x80\xf2\x2a\xba\xd2\x4c\x97\xea\xfc\x53\xb7\xd6\x98\x42\xf2\x70\x4d\x0e\x8b\x8d\xdb\x1e\xc3\x5f\x5e\xe8\x0b\x85\x92\xb3\x8c\xff\x89\x4d\xe5\x57\x0e\xd1\x14\x7d\x55\x88\x5c\x21\x4c\x25\xd7\x28\x17\x33\x71\x3f\xee\x79\x22\x4d\x2f\x2b\x1d\x4f\x30\xe3\xa6\xe9\x6e\x97\xa3\xf6\x6f\xbd\x66\xbe\x32\xa1\xf0\x63\x98\x0f\x8c\x4b\xe2\x4a\xe5\x9c\x67\x0e\xed\x3e\x17\x29\xd3\xb8\x16\xef\x4c\x1f\x46\x07\xea\xcf\x87\xb2\xb8\x28\x6f\x32\x9e\xfc\x97\xa7\x6d\x40\xbb\xf8\x7c\x4d\x3f\x5b\x30\xd6\xeb\x2d\x43\x32\xab\xe7\x7e\xb0\xcc\xca\x5a\x08\xce\xbe\xc0\xac\x0e\xc8\xae\xe1\xcc\x86\xe3\xec\xa4\x0b\xe2\xae\x2a\xe9\x23\x36\x0a\xc2\xe8\x14\xb5\x23\xf7\xeb\x98\xf9\x16\xcf\x7e\x10\x77\x9b\x15\xef\x40\x40\x3d\xb7\xe9\xf7\x3c\x87\x82\x49\x36\x51\x1b\xc3\xd7\xb3\xca\x51\x48\x13\xfc\x4a\xfd\xb3\xd4\x4a\x45\x8d\x12\x72\xa1\x61\x28\xca\xdc\x58\xf4\x48\x8e\xa0\x05\x35\xcf\x13\x31\xa1\xe5\x49\x85\x2d\x73\x68\x47\xb1\x23\x86\x27\x98\x08\x0a\xb7\x71\x95\x48\x67\x41\x18\x46\xf6\x59\xf0\x63\xc5\xe1\x45\x40\x97\x1a\x66\x75\xf0\x6f\x44\x3a\xdb\x1a\xe8\xe9\x0f\xa5\xb4\xdc\xad\x83\x3a\xad\x38\xc6\x50\x07\x7e\xc7\x78\x47\xd5\xbc\x75\xfe\xa6\x4c\x33\x3f\xae\xe3\xd1\xdd\x30\xa7\x57\x5a\x68\x55\x6f\x42\x39\xd7\xb6\x1b\x50\x69\xfa\x76\xb7\x99\x0d\xaf\x12\xd3\x82\x49\x85\x2e\xed\x69\x9d\x0b\xa5\xcc\xa2\x7f\xd3\x2e\x67\xf3\x2a\xf8\xa6\x4b\xbd\x45\x3c\xf7\xd2\xbf\x56\x19\xb6\xb4\xb1\x0d\x84\x43\xdc\xc5\xfe\x46\x42\xea\xc1\x60\xd9\x8c\x6a\x63\x76\x82\x19\xbe\x65\xff\x3b\xe9\xff\xd6\xbf\xee\xaf\x68\x7a\x56\xb9\xfd\x34\x3d\x2b\xeb\xc9\xa6\xf7\x04\xb3\xda\x9d\xbb\xed\x5e\x7f\xd9\xe6\x35\x9c\xe8\xcd\xb5\xd8\x1d\x0a\xef\x0c\x71\x9b\x8c\x5e\x86\xb8\x2f\xec\x9d\x26\xf5\x96\xee\x99\xfe\x06\xda\x37\x07\xda\x53\xd4\x6f\x87\xb2\xa7\xfd\xeb\x2d\x4e\x4d\x4e\x51\xef\x07\x73\x4f\x71\x71\x1f\xbf\xaf\x5d\x86\x8d\xc6\xdf\x20\xfd\x4a\x3b\x0c\xa3\xc2\xca\x2d\x46\xc7\x79\x1a\xd3\xe5\xa7\x47\x2e\xdf\x96\x22\xe2\xae\x4e\x91\xde\x19\x22\x3e\x7d\xb2\x54\xfb\x70\xdb\x26\x11\xa0\x94\xef\xcb\xd8\xef\x0e\xfe\xd7\x1d\x20\x9d\xa2\x3e\xcc\xb2\xb5\xd0\xee\xe6\xb0\x2c\xdb\xe4\x6e\xc2\x1d\x9e\xaf\x02\xf8\x17\xc1\xfa\x61\x96\xed\x0d\xd9\x0f\xb3\xec\xed\xc0\xdd\x86\x64\xc7\xf8\x7e\xcf\x24\x08\x99\xa2\xec\xda\x8f\xa3\xe6\x06\x8e\x0a\x5a\xa4\x4b\x81\x3f\xf0\x0d\xb1\x1f\xfe\x02\x0e\xea\x89\x56\xa6\xb2\xa2\x16\x69\x14\x58\x3e\x0d\x49\xc7\xcc\x81\x03\x43\x48\x95\x05\x98\x29\x9c\x1f\xf2\x99\x4a\x7c\x1a\xab\x21\x65\xbd\x06\x47\xb3\xe7\xe8\x70\x34\x5b\xa3\x85\x19\x6c\xd5\x56\xa3\x0d\xf9\x8a\x60\xcb\xa4\x99\xfd\x59\xb0\x11\x0e\x44\xb7\x3e\xbb\xbe\x40\x79\xc1\x46\x08\x3c\xd7\x56\x7d\x59\x14\x6b\x5b\xe7\xc2\xc4\x96\x35\x0b\x23\xb6\xed\x1c\x50\x74\x12\x91\xdf\x47\x87\x5a\xf0\x40\x16\x45\x68\x0d\x5f\x80\xd3\x6d\xfa\x8b\xa4\x11\xbc\xc7\xda\x08\x59\x59\x91\x97\x93\x1b\x94\x90\xa2\x66\x3c\xdb\xb8\x1b\x7f\x73\x8c\xd5\xb5\x4f\x37\x47\x36\x93\x9b\xd6\xed\x14\x9c\x25\xde\x80\x03\xf8\xf0\xcf\x3a\x47\x8a\xd1\xfa\xe5\xc9\xbc\x63\xab\x98\x2d\xf1\x67\x31\x7a\x45\x77\x16\xef\xc8\x85\x2d\xc3\x5b\xae\x6b\xad\x69\xd6\x5c\x8c\x35\x98\xdb\x5e\xdd\xcc\x83\x46\x77\x55\x25\x3c\x63\xf1\xb3\x4d\xb3\xdf\xc7\x45\x18\x47\x45\xc2\xc8\x9a\xb8\x32\xd3\xeb\x74\xec\x52\x4c\xc5\xcd\xaa\x86\x9e\x9a\xfb\xb2\x7a\xc8\x7a\xd5\x10\xcf\x79\x25\x5e\x74\x93\xc1\xe3\x7d\x2e\x8a\xbe\xf3\xa5\xcf\xd6\x2b\x1c\xc9\xf2\xf5\x7b\xd7\x91\xa1\x60\x49\x82\xca\x1e\xf1\x36\x63\x77\x98\x03\x1b\xd2\xae\x40\x29\x01\x99\x18\xf1\x7c\xd5\xca\xc6\x70\x99\xbb\x1b\x3b\xbf\xda\xec\x72\xcc\x68\xb8\xa7\x85\x0d\x89\x7a\xab\x75\x8d\x0b\xc4\x1e\x96\x35\xf5\x7d\xd0\xc3\x83\xf1\xa4\x64\x89\x8e\x2e\x58\x72\xc7\x46\xf8\xf8\x18\xb5\x9f\x7e\xbb\x84\xdd\xdb\xdd\xd1\x73\xaa\x78\xe9\xf5\xd0\x5f\xe3\x12\xa8\xaa\x5f\xb5\xa2\x17\x55\x55\xd2\x6e\x45\x8d\xeb\x77\xdb\x6c\xbe\x0b\x10\x95\x48\xb0\x48\x0e\x7c\x05\x48\x5d\x7b\x66\xb8\x7c\x87\xaf\xaa\xc3\xc2\x4b\x1c\x4a\x54\xe3\x27\x76\x95\x96\x86\x01\x7e\xe5\xca\xbc\x28\x78\x3b\xad\x51\xd8\x22\xaf\x79\xe5\x8c\x18\xd6\xe7\x43\xd5\x2c\x33\xbe\x0a\x89\x1d\xd1\x7a\x2c\x86\xc0\xfc\x6b\xfd\x10\xae\x40\x66\x67\xc8\x53\xd8\xbc\x0a\x9a\xb7\x41\x66\x27\x6a\x37\xd8\xfc\x74\xce\x55\x01\xdb\x27\x38\xbb\xdd\xe6\xbb\x43\xdb\xff\x0f\x3c\x6d\xe5\xf6\x5b\x20\xea\x7b\x04\xd0\x57\x84\xcc\x79\x34\xde\x16\x40\x0f\x9b\xd7\xb8\xd7\x5f\x6e\xb3\x36\x61\xa9\x50\xd6\x83\xf6\x75\xdd\x1a\x38\x89\xa3\x90\xfc\x4f\xfb\x5e\xf8\xd8\x00\xc8\x2a\xfc\x6c\x4b\xff\xf6\xbc\x6e\x19\x4e\xb6\x67\xec\x67\x21\xdb\x96\xf8\x24\x66\x12\xea\x90\xa7\x2c\x6e\xb6\x71\xc7\xa0\x29\x53\x0e\x50\xfd\x39\x2f\xf9\x5d\xdf\x0f\x6d\xd6\x37\x73\xed\x96\xf1\x14\xf5\xf2\x29\x61\xbd\xc1\x6c\xcf\x39\x00\xdf\x77\x07\x3b\x26\x03\xfa\x52\x0e\x84\x7b\x0f\xbf\x9e\x7b\x4d\x2d\x6d\xc3\x04\x7f\x19\x5a\xdb\xe8\xfb\x71\x4b\xc7\xed\x30\x85\xe6\x51\xb7\xe9\xda\x3e\x5c\xe3\x4b\x1d\xf4\x0b\x26\x15\xce\x59\x17\x34\xb2\x96\x9e\x0b\x2c\xd9\xa2\x57\x52\x88\xcc\x3f\x42\x26\x51\xfa\x0b\xb4\xf5\x3d\x8e\xc8\xb3\x19\xfc\x64\x89\x7e\x32\x33\x9b\x64\x67\x59\x26\xa6\x98\xee\xf2\xc2\x5d\x53\xeb\x8d\x6b\x8d\x8d\x37\xe7\x94\xf0\x63\xe7\xaa\xd7\xc1\xee\xc5\x7a\x6b\x03\xb8\x91\xb3\x6b\xf4\x5e\x62\xf0\x33\x17\xc9\x6b\x1c\xb5\x6f\xf8\x7f\x4e\x14\x5f\x71\x7d\x4d\x4c\xec\xd8\x8a\x9b\x78\x4b\x81\x3a\x58\xfe\x9f\x79\x5c\x07\x99\x3f\x01\xf9\x5f\x00\x00\x00\xff\xff\x79\x03\x7b\xfd\x6f\x34\x00\x00"),
          path: "jwt-api-http.tml",
          root: "jwt-api-http.tml",
        },
      
        "jwt-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x4f\x1b\x39\x10\x7f\x26\x52\xfe\x87\xb9\xe5\x05\x2a\x48\xf8\x6a\x0b\x91\xee\x81\x02\xe5\x52\xf5\x4a\x04\xa9\xfa\x80\x2a\xe2\xac\x27\x89\xcb\xc6\xde\xb3\x67\x03\x51\xb4\xff\xfb\x69\xbc\xbb\x49\x08\xe4\x03\x84\xee\xa3\x6a\x5e\x82\x9d\x99\xf1\x78\x7e\x33\xbf\x19\x33\x1a\x55\xae\xc8\x26\x21\x55\x2e\xda\x3f\x30\xa4\xca\x17\xd1\xc7\x34\x85\x3f\x9a\xcd\x06\x1c\x37\xea\x50\x2e\xfd\xbe\xf8\x53\x2e\x95\x4b\xd7\xbf\x5d\x9f\x1b\xb8\xc4\xd8\x58\x82\x13\x61\xe5\xf7\x8d\x1e\x51\xec\x6a\xd5\x6a\xd7\x58\xbf\x1d\x0a\x2b\x2b\xa1\xe9\x57\xdb\x42\x76\xb1\x3a\x1a\x55\x1a\x22\xbc\x15\x5d\x6c\x08\xea\xa5\xe9\xe6\x02\x8d\x6c\xf9\x58\x85\x4f\x5e\xea\xbf\x72\x20\x40\x24\x64\xa0\x8b\x1a\xad\x20\x94\xc0\x27\x81\x88\x15\x74\x8c\x05\xea\x21\x38\x6f\x02\x5a\x4f\x5a\xcb\x4d\xb6\x2a\x7c\x5e\xb3\x87\xde\x2c\xde\xc7\x18\x92\xf3\xda\x89\x43\x0b\x64\x40\xf5\xe3\x08\xfb\xa8\x09\x84\x96\x10\x5b\x33\x50\x12\xbd\x44\x5b\x84\xb7\xa8\x25\x28\x4d\x68\x3b\x22\x44\x16\xcf\x05\x24\x24\x5a\xa2\x8d\x94\x46\x90\xed\x4c\x44\x84\xa4\x8c\xae\xf1\x81\xad\x56\xab\x6b\xca\x25\x1a\xc6\x08\x75\x89\x9a\x14\x0d\x3f\x3c\x32\x37\x2a\x97\xd6\x4e\x4c\xa2\x69\x23\x34\x9a\xf0\x9e\x2a\x27\xd9\xf7\x26\x6c\x28\x4d\x5b\x80\xd6\x1a\xbb\x59\x2e\xad\x9d\x62\x84\x84\xb3\x62\x5b\x1c\x02\xa5\xbb\x9b\x99\x60\xb9\xb4\x76\x8e\x8f\x6c\x4d\x84\x36\x0a\x4f\xa6\x0c\x7f\x8d\xa5\x98\x6f\x78\x6b\xec\xfc\xe4\x88\x4b\x1c\x98\xdb\x15\x5c\xb9\xc4\x8e\x45\xd7\x5b\xe0\xce\xa7\x6f\xcd\xe3\x84\x7a\x53\xde\x9c\x23\x1d\x47\xd1\x7c\x6f\x8a\x6f\x1f\x1c\xa5\x39\x4e\xd7\xdf\x27\xb7\x7a\x18\xb3\x73\x2b\x1e\x47\x76\x0b\x46\x23\xff\x37\xa3\x55\x64\x66\x9a\x56\xa6\x77\x1f\x67\xd1\x93\xce\xf2\x92\x4f\x0e\x17\x04\x70\x13\x36\x26\xd9\xf9\xe0\xb4\xb9\x19\x3b\x39\x21\xf5\x79\xc4\xe9\x54\x27\x4e\x5d\xe3\x30\x4b\xdd\x8e\x89\x22\x73\xa7\x74\x17\xfa\x48\x3d\x23\x9d\x2f\x09\x14\x61\x0f\x50\xcb\xd8\x28\x4d\x3e\x0b\xd7\xd7\xc1\x07\x01\x46\xd0\xb8\xb8\x6a\x42\x75\x4e\xdd\x55\x21\x65\xe1\x75\xf8\xd3\x9b\xab\x41\xab\x93\xe8\x10\x36\xb8\xd8\xde\xf0\xc5\x1b\xf5\x4d\xc8\xc3\x49\xf7\xf0\x86\x2b\x31\x21\x15\x4d\xd2\xd5\xbb\xec\x5d\xbd\xc4\xbf\x12\x74\x04\x91\xe9\x2a\xcd\x25\x5c\xc4\x48\x19\x0d\x89\x63\xa7\x9f\x44\x60\x0e\x00\x69\x0a\xbe\x88\xee\x7a\x2a\xec\x31\x2d\x48\x8c\xd4\x00\x2d\x4a\x10\xae\x5c\xfa\x74\x75\xf1\x85\xab\x92\xb2\x02\xaf\x40\xb3\xa7\x1c\xdc\xa9\x28\x02\xa5\x81\x12\xab\xc1\xa2\xff\x12\x60\xd1\x71\xf5\xab\x01\x33\x87\xa0\xc4\x41\x68\x24\x7a\x7e\xd8\x86\x33\xcf\x0c\x28\xa1\xf0\xff\x83\x91\xc3\xbc\x92\xf9\xba\xe5\x12\x00\x80\xbf\xaf\xa6\xed\xe6\x30\xc6\x1a\x88\x38\x8e\xf2\x9b\x55\x7f\x38\xa3\xc7\x70\xb5\x5a\xad\x6c\x3d\x1a\x41\x5f\xc4\x97\x42\x4b\xd3\xf7\xae\x8e\xef\x08\x01\x4b\x04\x10\x04\x90\x4e\x70\x9e\xf2\xe3\x2a\x73\xf1\xc4\x48\xcc\x4d\x96\x4b\x57\x49\x18\xa2\x73\x35\xd8\xdb\xd9\x79\x4a\xe7\x12\x5d\x6c\xb4\xc3\x57\x72\xbe\x5c\x5a\x0b\xf0\x3e\x56\x16\x5d\x50\x83\xfd\xbd\xfd\xbd\xfd\xdd\x2d\xde\xb4\x59\x69\xdf\x90\xb9\x45\x1d\xd4\x82\xc0\xef\x0a\xef\xdc\xcc\xa6\x5f\xdd\x30\x84\x41\x0d\x82\x0f\x28\x2c\xda\xe0\x81\x91\x27\x4e\x98\x04\x64\x7d\x7d\x49\xe6\x5a\xcf\x46\xcb\xb3\xb7\x60\xad\xe5\xe9\xcb\x72\xdc\x7d\xba\x6a\x80\x1a\x3e\x7d\x6b\x42\xee\x29\xf8\xbb\xcc\xe6\xcb\x53\x38\x7d\x14\x2a\x4a\x2c\xd6\x58\xf0\x2d\x43\xb5\x0d\x07\x3b\x07\xfc\xf3\xcb\x11\x7c\x2e\x7a\xef\x3b\x72\xf7\xed\xd1\xfe\x61\x78\xb8\xb7\x1f\x76\xde\x1e\xe2\xfb\xc3\x9d\xdd\xa3\x36\x8a\x3d\xd1\xd9\x3d\xd8\xeb\x1c\x1d\x1c\x1c\xbd\x3b\x7a\x27\xa6\x43\x7d\x7e\xb6\x30\xd2\x3e\x0a\xab\x84\x3a\x27\xfe\xe5\xb1\xf6\x82\x33\xd1\xce\xd2\x28\x0b\x36\xb4\x87\xe0\x50\x4b\x66\x0d\x6d\xf4\x76\x96\x2b\xb2\x40\x04\x65\x2e\x26\x1c\xf0\xad\x73\xd2\x7d\x1e\x44\x6b\x39\x46\x6b\xff\x13\x90\xea\x5f\x3e\x5e\xcc\x47\x69\x39\x3c\x75\xdd\x31\xcb\xb1\x61\xa9\x9c\x3a\x19\x0b\x12\x11\x98\x0e\x58\x0c\x8d\x95\x0e\xc4\x40\xa8\x48\xb4\x23\x64\x96\x1d\x8f\x64\xcc\xd3\xcf\xed\x76\xcf\xc1\x2a\xc3\xe9\x1f\x27\x41\x56\x0b\x7c\x08\x82\x1a\xec\xee\xcc\xf2\xd3\xc2\xa2\xa9\xc5\x49\x3b\x52\xe1\x8d\x92\x2b\x34\x58\x5c\xa1\xbd\x9e\x23\x31\x2c\x56\x0d\x8a\xb2\xe1\xda\xc8\x70\x61\x84\xb8\x11\xce\x9b\xae\x3d\x40\x1d\x6b\xfa\x5e\x6a\x3c\x6b\x67\x28\xe7\x66\x50\x0d\x7c\x7d\xb9\x24\x22\x2e\x2b\x01\xdc\xb7\xca\x25\x9b\x47\xb5\x02\x75\xe2\xd1\x39\x1b\x44\xc6\xc3\x70\x6b\x72\xd1\x16\xc4\xc2\x8a\x3e\x12\x5a\x36\xe0\xc5\x78\x83\xf2\x71\xdb\x4f\x6b\x9d\xa1\xff\x21\xf3\x7b\x5e\x0b\x6e\x14\x76\xdc\x38\x15\x18\x8b\xe9\x98\xae\xd6\x37\xff\x9d\x04\x7a\xf4\x00\xc8\x1f\x2c\xdc\x57\x1b\xfe\x0a\xf5\x53\xc8\x3f\x19\x71\x15\x2b\x9f\x7b\xb5\x60\x7c\xcf\xa0\x35\x99\xa7\x9b\x9e\xf0\xe6\xa9\x3c\xec\xcc\xac\x76\x96\xb5\xd7\x42\x0e\x48\xf5\xb1\xd2\x54\x7d\x9c\x56\x2b\x7a\x30\x2b\x34\x85\xed\x22\x2d\x71\x8d\xbc\x50\xe1\xda\x67\xe1\xe8\xb3\x9f\xf7\x16\x9c\x11\x09\x47\x37\x7e\x2a\xf4\x3a\x75\xe7\x12\x94\xc7\xb4\xd0\xaf\x19\x9d\x3c\x04\x75\x7e\x40\x0d\x44\x94\xe9\x9c\x26\x36\x1b\x2e\x67\x42\xa0\x72\x29\xaf\x79\x2a\x48\x3c\x9b\x9b\x0a\x8b\x52\x90\x60\x2b\xab\x97\xfd\x4a\xc5\xee\x5f\x3a\xcf\xab\xf7\x28\x1a\x53\xf0\x8b\x6a\xfd\xbf\xce\xb6\xd7\x45\xa9\x8c\xd2\xc9\xd3\x73\x94\x7e\x7f\x30\x0f\x7e\x7d\x25\xbe\x2d\x9e\xbe\xcb\x20\xc8\xe4\x40\x10\x61\x3f\x26\xee\x86\x90\xe4\x5b\x3f\x07\xff\xc2\x9d\xa2\xde\x43\x7b\xfe\xb1\x62\x73\x1e\x6e\x1b\x39\x7c\x55\x8e\x7e\xa5\x37\xd6\x2f\x82\xfd\x69\x09\xf6\xa5\x1c\x75\x30\xcd\x14\xa7\x67\x9f\xcf\x9a\x67\xaf\x43\x16\xc5\x3f\xe0\x56\xe1\x6b\xe9\x65\x7f\x8d\x67\xaf\x88\xe6\xdf\x01\x00\x00\xff\xff\x35\x7c\x1d\xc0\xf1\x16\x00\x00"),
          path: "jwt-api-readme.tml",
          root: "jwt-api-readme.tml",
        },
      
        "jwt-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3b\x5b\x73\x1b\x37\x77\xcf\xcb\x5f\x81\xf0\xc1\x25\x3f\x33\xab\xb4\xe3\xc9\x03\x53\xcd\xd4\xb2\x9d\x44\x99\xc4\xf1\x58\x4a\x32\x9d\x8c\xc7\x86\x76\x0f\x49\x48\xcb\x05\x0b\x80\xa2\x55\x55\xff\xbd\x73\x0e\x2e\x0b\xec\x85\xba\xd9\x6e\xfa\x4d\xfc\x90\x48\x8b\x83\x83\x73\xbf\x01\x1a\x1d\x1c\x30\x50\x4a\x2a\xcd\xf2\x3c\x1f\x5d\x72\xc5\x26\xa3\xec\x95\x52\xaf\xa5\xf9\x5e\x6e\xeb\x92\x1d\xba\xf5\xfc\x35\xec\x26\xe3\x5a\x1a\xb6\xc0\xef\xe3\x29\x81\x1d\xd7\x97\xbc\x12\xe5\x71\x09\xb5\x11\xe6\x8a\xb1\x16\xfc\x46\xc9\x4b\x51\x42\xc9\x02\x84\xd0\x4c\xd8\x4d\x0e\xc5\x1b\xae\x34\xfc\xf4\xc7\xe9\xa9\xbc\x80\x9a\xb9\x7f\x2d\x2c\x08\x62\xbf\xcc\xfd\x6e\x76\xbe\x33\xcc\xe0\x1e\x87\xe7\xb5\xfc\xe9\x8f\xd3\xe7\x5b\xb3\x92\x4a\xfc\x37\x37\x42\xd6\x16\x63\x9b\x01\xda\xc8\x63\xb8\x04\xcd\x71\x6d\x40\xd5\xbc\x7a\x85\x9b\xfa\xc9\x39\xe3\xc5\x05\xd4\x25\x5b\x70\x51\x41\xc9\x76\xc2\xac\xec\x7a\x2a\x94\x94\xa7\x14\x85\x82\x02\xc4\x25\x44\x5c\x74\x05\xf3\x5b\x0d\x1f\x37\x50\x18\x40\x4c\x2f\x2a\x2e\xd6\x1d\x34\xcd\xee\x82\xd6\x85\x66\xa8\x21\xbf\x8f\x99\xab\x0d\x38\x6c\xaf\x3e\x6e\x84\x82\x07\x10\x05\x76\x63\xca\xdb\x5b\x58\x28\xd0\x2b\x87\x2a\x45\xe3\x40\x98\xb2\x30\x89\x74\x69\x83\xdb\xfc\x12\x6a\x01\x65\x0f\x15\xb4\xfa\xde\x12\xc0\x2b\x05\xbc\xbc\x62\x67\x15\x2f\x2e\x2a\xa1\x4d\xca\xce\x3e\x3a\xfa\x11\xa5\xec\x38\x34\xcf\x8b\x02\xb4\xee\xc5\xc2\x69\x69\x2f\x12\xc7\xf0\x89\x58\xd6\xa2\x5e\xfe\x02\x66\x25\xdb\x7e\x63\xb7\x6b\x0b\xc1\xd6\x16\x64\x2d\xf4\x9a\x9b\x62\x45\x88\xa6\x23\xf4\xc5\xe0\x26\xb0\x3e\x93\xa5\x00\xcd\x4a\x6e\x38\xd3\x46\x2a\x28\xd9\x42\x2a\xc6\xd9\x56\x83\xfa\x17\xcd\x2a\xb9\x14\x35\x2b\x14\xd0\x16\x5e\xe9\x7c\x84\xea\x6e\x50\x68\xa3\xb6\x85\x61\xd7\x23\xd4\xf4\x9b\xed\x59\x25\x8a\xe3\x97\xde\xa2\xb5\x51\x48\x89\xfb\xf7\xe1\x5c\xcb\x7a\x3e\xde\x10\xd0\x7b\x51\x8e\x3f\xd0\xa6\x54\xba\x03\x9b\x12\x31\xbb\x8d\x56\xac\xda\x43\x32\x23\xd6\x90\x9f\x8a\x35\xc4\x1b\xad\x10\xb5\xdb\x72\xca\xd5\x12\xcc\x2d\x04\x1a\x02\x6a\x08\xfc\x99\x6b\xf3\x33\xc9\x61\xcf\x39\x15\xd7\xe6\x3d\x49\xcb\xed\x3a\xd6\x7a\x0b\xe5\x73\xb3\x97\xba\xce\x2e\x27\x0c\x8a\x0f\x97\xbc\xb2\xbb\x5e\x6e\x95\x0d\x21\x2d\x61\x08\x07\x35\xfe\x30\xca\x5e\xa2\x06\xaf\xaf\xf3\x13\xd2\x47\xfe\x86\x17\x17\x7c\x09\x37\x37\x79\xf3\xed\xd7\xb3\x73\x28\x4c\xfe\x9a\xaf\x81\xfe\x73\x73\xe3\x11\xa2\xfa\xc7\x1f\x46\x37\x64\x1d\xbf\xa3\x99\x71\x03\x4c\x81\xd9\xaa\x5a\x33\x5e\x5b\x2b\x63\x62\xc1\x96\xe2\x12\x05\x16\xd4\x5f\x4a\xb0\xd1\x80\x6c\x8c\x95\xa0\xd1\x64\x11\xcd\x42\x40\x55\xb2\x4b\x5e\x6d\x81\x69\xc3\x0d\xe4\xa3\xc5\xb6\x2e\xd8\x44\x34\x71\x7a\x1a\x0e\x9b\x4c\xdd\x19\xd6\x90\xc4\xc2\xa9\x46\xe7\xa7\x4a\xac\x4f\x36\xbc\x80\x89\x28\x73\x6f\x60\x53\x76\x78\xc8\xc6\x63\x07\x8d\xff\x2c\xb1\x89\x3b\x04\x6b\x14\x9a\x29\xf8\xaf\xad\x73\x26\x84\xbe\x19\xed\x3d\x26\x36\xc9\x3b\x1d\x95\xd8\xf0\xbd\x8f\xf3\x56\x79\xa7\xa3\x82\x09\x0f\x1e\xe3\x36\xd5\xa2\x72\x1a\x75\x29\xab\xe5\xee\x21\x6b\x72\x44\xa4\x37\xb2\xd6\xc0\x8c\x64\xdc\x45\x65\x1f\x59\xa5\x62\x5a\x4b\xc4\x63\x83\x81\xdc\x80\x35\x47\x17\x0a\x3c\xf6\x24\x12\x24\xb1\xae\x71\x34\x67\x6f\x71\xb8\xdb\x1f\x05\x06\xdc\x3f\x23\xb8\x53\x3c\x9e\xf5\xe0\x27\xb0\xf7\x94\x97\xfa\x43\x85\xa8\xcd\xb7\xcf\x06\x62\x84\xa3\x24\x6c\x49\x61\x3d\x25\xcd\x9e\xbb\x7b\x8d\x97\xd4\x03\x9c\xe6\x9c\xfb\xdd\xbd\x3e\x93\xf5\x9a\xd6\x39\xcf\x83\x9c\x1a\xdb\xca\xb2\x3e\xab\x0a\xf2\x4c\xcd\x2a\xbb\x19\x0d\x23\x8f\x94\x7c\x0b\xfa\xd8\x1c\xee\x71\x40\xbf\x1f\xf6\x9e\xb0\xc7\x03\xfd\x11\xe7\x3c\xf7\x4a\xfd\xf7\x43\xf6\xcd\x10\x26\xd2\xac\x2b\xd9\x30\x5c\x8b\x9a\x6d\x6b\xf1\x71\x18\x69\xcb\x60\xf6\xe1\xf6\x3e\x75\xe7\x33\xb2\x3e\x67\xb6\x55\x63\x2b\x79\x43\x6d\xd0\x93\xad\x45\xd8\x1c\xde\xe3\xa9\x76\x6b\xe2\xaa\xaf\x14\x55\xa0\x76\xa3\xf7\x08\xa5\x9c\x37\x9c\xa8\x02\x01\x92\x55\xad\x0a\x0b\x60\x29\xb2\x38\xbd\xdd\x9b\x15\xb0\x6d\x5d\x82\xaa\x44\x0d\xac\x90\xeb\x33\x51\x43\xc9\xb4\x2a\x18\xaf\x4b\x44\xe4\x90\x91\x95\x6b\x44\xc0\xb5\x96\x85\xe0\xc6\x97\xb9\x88\xc2\xf9\x4e\xad\x0d\xaf\x8b\xe0\x06\x26\xf0\x30\xb5\xc7\x4e\xa6\xde\xf5\xaf\xe3\xd8\x67\xf2\x57\x4a\xe5\x1e\xe2\x29\x1b\xb3\xc9\x98\x3d\x65\x26\xb7\xec\xc4\x2b\xd3\xb1\x63\xc3\x27\xa5\x97\x47\x47\xae\xf4\x2e\x61\x21\x6a\x20\x57\xa6\x54\xbb\xe0\x05\xb0\xdd\x4a\x14\xa4\x40\xa9\x71\x29\xe2\x15\xeb\x27\xbe\x04\xa6\xaf\xb4\x81\x35\xb9\xb3\x95\x8b\x12\x40\xfe\x8f\xfc\x23\x10\xfe\x2c\x7c\x06\x55\x50\x48\x55\xea\x1c\xe1\xff\x43\xac\x37\x15\xac\xa1\x36\x69\x95\xd5\x90\xd4\xd0\x71\x3d\xca\x5e\xc8\x6d\x6d\x26\x85\xf9\xc8\x0a\x59\x1b\xf8\x68\xf2\x17\xf6\xff\x53\x36\x11\xb5\x99\x59\x21\x4e\x47\xd9\x4b\xa8\xc0\x40\x1f\xe4\x8c\x6d\x7c\x96\xb4\x72\x74\x31\x65\x94\xbd\x50\xc0\x87\xf6\x40\x05\xeb\x28\x89\xbb\x1d\x3f\x40\x2f\x31\x3d\x47\x4c\xfc\xde\x86\xc4\xdf\x36\xe5\xe0\x71\xad\xfd\x7b\xce\x7f\x5e\x55\x47\x57\xbf\xaa\x12\x54\x3f\x26\x89\x4b\x01\x0d\xfd\x76\x74\xd5\x90\xf5\xe7\xbb\x2e\x61\x3f\x80\x39\xba\xfa\x1e\xe3\x72\x3f\xca\x0b\xb8\x0a\x08\x6d\xe0\x0e\x3a\xba\xbe\xe9\x65\xd5\xd2\x79\x7f\x02\x67\x6c\x83\xe6\x45\x9a\xf5\x09\xfb\x0d\xa8\x37\xee\x63\x8b\xfe\xd8\x00\x52\x03\xff\x75\xa3\x9b\xe0\xe1\xba\x04\xbd\x81\x42\x2c\x44\x41\x26\xbb\x54\x1c\x37\x63\xfb\x8a\x3b\x0a\x6e\xe8\x74\x05\x97\xf2\xc2\xdb\xb1\x8b\x64\x91\x29\x0b\x68\x37\x07\x78\x50\x62\xb0\x6f\x11\x03\x4c\x3a\x5c\xb7\x4c\xcf\x05\xd4\x61\xb8\x89\xcb\x85\x91\x4c\x91\xe4\xee\x86\xeb\x6b\xfa\x59\xf1\x76\x49\x1c\xbe\x76\x8b\xe2\x3e\xf4\xcf\x1b\x51\xec\x21\x7f\x72\xdf\xfa\x7b\x40\x41\x7b\xe3\x8f\x59\x71\xe3\xc2\x0f\xe3\xec\xac\x13\x17\x6c\x7c\x2a\x78\x4d\x53\x17\x0b\x67\x95\xac\xed\x5e\x24\x9f\x0b\x6c\x33\x2b\x56\x03\xe6\x65\xae\xae\xa8\xb6\xb3\xda\x27\x54\xbc\xa0\xe4\x44\xe1\x98\x6f\x04\x2d\xac\x8c\xd9\x30\xa8\xcb\x8d\x14\xb5\x69\xeb\xba\x37\x40\x45\x86\x10\xa2\xd5\xdd\x22\xd5\x2d\x16\x42\xa1\x66\x50\x0d\xc3\xd1\xa5\x7f\xcb\xac\x13\x4b\x30\x95\x78\x37\x1d\xda\xe3\xff\x4f\xf4\xdf\xc7\xff\x7e\x3c\x3d\x7d\x73\x7b\x6e\x49\x57\x16\x52\x21\x0a\x52\x81\xd7\x87\xc7\xd7\xd2\x04\xa1\x4f\xd4\x80\xd2\xfa\x07\x6e\xdd\x1a\x51\x35\x92\x77\xb2\x3c\xae\x17\x72\xcf\xb2\x13\xdd\x30\x80\x53\xd9\x30\x80\x13\xe4\x1e\x00\x72\xde\xe1\x75\x17\x36\xf6\x01\xd8\x80\x31\x0c\x91\x38\xf0\x20\x58\xaa\x27\x14\x8c\xf7\x16\x0a\x94\xfc\x6b\x2a\xb5\x14\x2c\xb9\x2a\x29\x10\x56\x95\xcf\xdf\x58\xc2\x95\x67\x4c\x2e\x58\xa2\x8c\x96\x6e\x08\x65\xa8\xbe\x32\x76\x2a\x0d\xaf\x50\x59\x4d\xe7\x62\xa8\x9f\x4f\x28\xa1\x22\xcf\xdb\x8b\xc5\x6f\x24\x53\xb0\x51\x40\x65\x1f\xd6\x4b\xa1\x83\x5b\x8a\x4b\xa8\x6d\x1f\x87\xa5\x24\x68\xe3\x8d\x27\xa6\x16\xe9\x5c\x81\xc5\xd5\x17\xb5\xfa\xe9\x47\x42\x22\xea\x7d\xe6\x09\x93\x1d\xbe\xc4\x86\x2b\x63\x6f\xbb\xd9\xa9\xe9\x9f\x92\x25\x02\x27\x29\xbc\x0d\x72\x4c\x85\xf1\xde\x91\xec\x10\x5b\xa0\xc6\xd1\x1a\xbc\x1e\xca\x8a\xee\x14\xb4\x11\x6b\x59\xb7\x87\x5c\xb6\x40\x84\x92\x9d\x5d\xd1\x8c\x8b\x3c\xa9\xc0\x5a\x07\xe3\x9d\x5c\xb0\x1a\x76\x4d\x71\x46\x03\x4f\x27\x83\x06\x63\x52\x3e\xc7\x23\x25\x1b\x12\xe8\xf3\x83\xe6\x31\xa3\x9b\xd0\xf9\xbb\x52\xdf\x4e\x64\x03\x0b\xa8\x33\x62\x83\x1b\xc3\x8b\x55\x5c\x30\x63\xa1\x5c\x72\x55\x5a\x9a\x75\x1e\x10\xf9\xd2\xdf\xa2\x4a\x68\x3f\xdf\x99\xfc\xc4\xed\xa3\x65\x3d\xca\x4e\xb0\x12\xe0\x55\xa8\xb6\x42\xd1\x6f\xbf\x37\x63\xb1\x87\x72\x18\xbb\xd8\x2f\xfc\x02\x54\x64\xdb\x58\xe5\xdb\xae\x08\x89\x0e\x13\x8a\xb3\x2b\xb6\x26\x48\xaa\x4e\xa0\xa6\xb6\x26\x2e\xa2\x63\x3d\xa5\xb8\x11\x63\x37\x86\xa3\x38\x64\xbd\x10\xcb\x07\x97\x08\xc1\x1a\x42\x90\x27\xab\xc3\xf6\xd3\xf5\xe7\xf2\x6e\x8c\x05\x23\xbc\xb4\xdb\x88\xaf\xba\x90\x6b\xfc\x81\x06\x19\xde\xfe\x52\xdc\xc4\x58\xc4\xc7\x3f\x50\x99\xae\x53\x9e\x44\x25\x68\x42\x5f\x80\x4f\x4d\xca\x8e\x1d\xc8\x13\x68\x75\xab\xdc\xb8\xc2\x65\xf5\xa6\x7f\x74\xbb\x13\x2b\x3a\x11\xcb\x1a\xc2\x2d\x46\x3c\x8f\x09\xcb\xa2\x5e\x9e\x40\xa1\xc0\x74\x97\xad\x9a\xd2\x7f\x89\x0a\x2d\x94\x2d\x53\x93\x7f\x64\xbd\xf1\x38\xbc\x3d\x79\x0a\x23\x9c\x64\x8a\xda\x19\x37\x79\xb0\x2e\x54\x94\x35\xc2\x35\x4d\xaa\x86\xa6\x29\x6f\xee\x9e\x7c\x0f\x67\x65\x3b\x5c\x20\x51\x3b\xec\xd5\xc0\xab\x0a\x11\x35\x05\x99\x2b\xcc\x43\x3d\xbe\xb7\x04\xa7\xdb\x26\x62\x3c\x5e\xb1\xe6\xa3\x1b\xed\x75\xa7\xf6\x99\x55\x38\x32\x16\xb4\x3b\xca\xc2\x3d\x88\xc7\x40\x4c\x9f\x80\x19\x65\xbb\x95\x30\xd0\xbf\xd4\xe9\x56\x9d\x70\x5e\xc3\x2e\x3e\x3b\x0c\xc8\x28\xd4\x26\x92\x73\xdd\xbe\xaf\x83\x82\x6c\x50\x6c\x7b\xa5\x8a\xc7\x04\xc1\xba\x61\x41\x7a\xec\xc4\x31\x1a\x79\xcc\x20\x2b\xb3\x41\xfe\x67\xa1\xe0\xee\x30\x3b\x4d\x58\xb9\x0e\x93\x9b\xe8\xeb\xf5\x28\x73\xf2\x9e\x3b\x47\x9b\x8d\xb2\x46\xd8\xf3\xe6\xfe\x09\xbf\x07\xf2\xe6\x0d\xa5\xf8\xbd\x73\xf2\xdc\x13\x35\x1b\x65\x3e\xc0\xc6\xf5\x0e\x26\x0b\x58\x6f\x50\x72\x32\x6e\xea\x80\x62\x8f\xf6\x86\x63\xa7\xbe\x46\xfa\x40\x04\xee\x0e\xc8\x4f\x20\x77\x3c\xe6\x65\xca\xd2\x9e\xa8\xaf\x99\xe5\xd1\xd4\xef\xd1\x3d\x12\x8a\xf4\x92\x2b\x77\x07\xe9\xf3\xd9\x68\x94\x11\xdd\x04\xc5\xe6\x87\x14\x14\xe8\xb2\xf7\x0f\x61\x56\x36\xa7\x4d\x22\x3a\x66\xec\x09\x21\x98\xd9\xf8\x69\x6e\x0d\x9b\x34\xc9\x13\x0b\x66\x72\x17\x82\xbe\xc2\x33\x78\x6e\xf5\xe7\x3f\x22\x50\x34\xaa\x9b\xb1\x81\x0b\xbb\x51\x46\x73\x43\x1a\x1c\x36\x48\xd2\x30\xf3\xd5\x21\xa2\x48\x50\x0e\xc1\x4e\x9a\x85\x19\x33\x53\x8f\xdd\xed\xfa\xf3\xdd\xd9\x95\x81\x08\x26\x4f\x42\xf1\x74\x46\x53\xc5\xec\x66\x6a\x27\x99\x28\xc0\xe8\x6c\xe4\x99\xa2\xa3\x52\x33\x26\x2f\x50\xb6\xa0\x54\x3e\x21\x81\xb9\xf8\x27\x64\x6d\x87\x71\xdf\x21\x04\x51\x1c\x6d\xb3\x53\x36\xfd\xa4\x67\xc3\x2f\xbc\x5a\x48\xb5\x06\x12\xa7\x9d\x96\x06\xb2\xef\x6b\x20\xa8\x2b\x3f\x15\xb4\x88\x32\x3b\xe5\x9b\x37\x0c\xd8\xcf\xf4\xcd\x4e\x3c\x5b\xb7\xe7\x16\xe0\x66\x94\x39\x09\x66\x3e\x97\xd3\xdd\xb4\x30\x2b\x50\xfe\x3a\x96\x49\x45\xb3\x7b\xec\x90\x2f\x81\x5d\x61\xe4\xeb\x65\x7b\xd2\xc3\xb7\xbb\x09\xfe\x9f\x9e\xa5\xd7\xd2\xd0\x97\xff\x04\x33\xfd\xbf\x92\x4a\xeb\xfa\x3e\x91\x4a\x6a\x5b\x8f\x26\xc7\x13\x03\x8e\x8e\x88\x8a\xe4\xad\xc6\x8c\x0e\xf6\xd3\xf6\xaf\x6c\x41\x44\x82\x8a\x67\xec\x8f\x26\x67\xdf\xf1\x31\xad\x3d\x8f\x22\xdc\x73\x08\xf7\xa8\x22\xa2\xf7\xe0\x80\x15\x2b\x28\x2e\x98\xa8\xbf\x5e\xc3\x5a\xaa\x2b\x56\x60\xd1\x9e\x53\x6d\xb1\xe2\x3a\x0a\x5a\x3c\x0f\x31\x3e\xff\x91\x6b\x8c\xa7\x33\x57\xd5\x1e\x97\xfe\xa7\x50\x9a\x4f\xfb\x1c\xf6\xaf\x2c\x0a\x54\xdd\x8a\xeb\xcf\x45\x67\xf7\x81\xc5\x63\x88\x75\x04\x5a\x99\x63\x97\x33\x8b\xae\x5f\xec\x2c\x82\xc1\x47\xa1\x8d\xa6\xea\x0c\xcb\x91\xcd\x56\xd1\xcc\x06\x13\x2c\xc5\x85\xea\xca\xd6\x6a\xc0\x38\xdd\x04\x22\xb8\x2f\xd4\x6c\xab\x9a\x16\x69\x88\xda\x56\x3d\x18\xde\xdd\x5e\xaa\x0a\x99\xef\xf2\xb1\x0a\xda\x40\x5d\x62\xe1\x3c\x90\x92\xfd\x7c\xb5\x77\x66\x5f\x17\xb2\x6c\xbd\x65\x49\xa6\x6a\xa8\x9c\x6d\x7d\xc6\x35\x94\xdf\x3e\x3b\x4d\x92\x2a\x7e\xfc\xf6\x59\x7e\x62\xca\x57\x88\x45\xd4\xcb\xfc\x25\x20\xba\x13\x42\x30\xe9\xc1\xbd\xcf\x44\xef\x1a\x08\x7a\x9e\x01\x45\x6a\x2a\xc1\x1f\x6a\x99\x99\x1f\x86\x4b\xc4\x93\x4d\x25\xcc\xc4\xfe\x36\x49\x79\x9a\xce\xd8\x78\x3e\xb6\xd4\x55\x50\x4f\x5a\x58\x28\xf4\xfe\xdb\x10\xad\xb7\xd3\xd6\x6b\x73\xe9\x1b\xa1\xd6\xed\xef\x8e\xd7\xc6\xbe\xba\x59\x73\xd3\xb2\xc3\x08\x37\xd3\xa1\x2f\x9f\x1f\xb2\x16\xd5\x7f\x7e\xf3\x6e\xd6\xf9\xf6\xaf\xef\x6c\x0c\x7a\x6b\x2f\xa0\x20\xb6\xbd\x85\x92\x6b\x5f\x35\xe6\x78\x12\x7e\x4d\xc2\x51\xa7\xca\xcc\xd3\x8b\x90\x19\x6b\x5d\xca\xcf\x98\xfa\xa2\xea\xc7\x22\xb7\x6c\x79\x11\xf6\x00\xb2\x09\xb7\x44\xc3\x7b\xcb\x96\xe5\x2a\x14\xd8\xf9\xf3\xd2\xb1\x61\x79\x0f\x2f\x4b\xc2\x87\xc4\x98\xbf\x7b\x34\x2b\xd1\x7b\xbf\x34\x2e\x36\xc4\x75\x45\xde\xdc\xdb\x75\xe8\xfc\x6c\x24\x1d\x1c\x30\x05\x6b\x79\x09\xde\x56\xec\x73\x28\x1b\x7e\xe8\x70\xba\xf2\x32\x92\x15\x5b\xa5\x28\x30\xd9\x4a\x08\xc5\xaa\xf3\x16\x4f\x4d\x56\x7b\x4b\x48\x07\x64\x1e\x8c\x3b\x65\xeb\xc9\x13\xff\x9b\xeb\xc3\xe2\x37\xa3\x77\x72\xd2\x16\x8f\x6d\x81\xa4\x11\x3f\x8e\xf2\xd6\xaa\xe2\xce\xc9\x0d\x6d\x5c\xdf\x9a\x36\x4d\x2b\x25\xb7\x4b\x3b\x87\x0b\xb3\x9d\x96\x65\x2e\x6c\xa6\xa1\x7b\xde\xe7\x55\x25\x77\xf6\xa6\x5c\x83\x42\xe4\x28\x61\x44\x6b\x81\xe9\xfa\xde\x1f\x10\xdd\x30\x2b\xd0\x72\xab\x0a\xd0\xf9\x70\x06\x70\x57\x67\xf7\x4d\x01\xed\x3b\xaf\x2f\x9a\x0d\xf0\xe0\xfb\x54\x88\x7f\x91\xc4\xd0\x4b\xf6\xdf\x39\xe2\x13\xe5\x88\x4f\x6b\x14\xa1\x18\xa7\x3b\x50\x7f\x0d\xd2\xf2\x51\x2b\x73\x51\x37\x83\x98\x7c\x94\xd1\x4b\xf4\x84\xef\x66\x35\x94\xea\x77\xca\x22\x9f\x9a\xdf\xc1\x00\x7e\xbc\x48\x98\x98\xf9\xa2\x92\x8e\x2a\xed\x8b\x68\x32\x3e\x1b\xae\x17\x3d\xf1\xf4\x36\x6a\x3a\x4f\x9a\x2c\x5a\x6b\x9c\x31\x99\xdd\xc7\xd8\x11\xad\xb5\xdc\xa1\x50\x69\xf2\xfa\x5a\xee\x26\x56\x44\x4e\x76\x6e\x2e\x9b\x1f\xc1\x42\x2a\x98\xd4\x72\x17\xa6\x31\x7b\x8d\xf0\x5e\x59\xf3\x36\x96\xf7\xa7\x93\x8e\x7e\x42\x9f\x7c\x70\xc0\x8e\x50\xfe\x94\x16\x53\x3b\xcb\x2d\x0f\xef\x07\x6c\xea\x91\x95\xc9\x67\x64\xe8\x71\x75\xc1\x17\x2d\x0c\x3e\x97\x14\x1e\xe1\x20\xd8\xff\xba\x19\x12\x32\x6a\x5f\x7d\x77\xdc\xa5\xe7\x4f\x0e\x22\x7f\xb1\x45\x81\xbf\xb0\x98\x1f\xb2\x5a\xee\xc8\x60\xe2\x09\x61\xe7\x02\x64\x3a\xea\x1d\x9d\x66\x4d\x87\xcb\x0e\xbd\xe8\xf1\x37\xbf\x72\x5c\x36\xdf\xbd\x4a\xc2\x9a\x7f\xde\x6e\x69\xf8\xad\x16\x1f\xd1\x7b\xdd\xa0\x62\x4b\xed\x7b\xb3\xdb\x5f\x92\x7a\x00\x47\x18\x6d\x4f\x78\x6a\x23\x0a\x09\xed\x90\x6d\xb7\xa2\x44\x89\xfe\xfe\x6c\x32\xcd\x5d\xa9\x31\xb5\xf1\x8e\x53\x23\xf0\xd0\x6a\x35\xe9\x10\x1a\xbb\x1c\xf6\xc3\xf6\x34\x66\x4f\x29\x7e\x7b\x89\x70\xb7\x0a\xf5\x7c\x67\x42\x31\x73\xbe\x33\x28\x87\x68\xc0\xdd\x19\x48\x3b\x1a\xa7\xb4\x31\x1e\x8b\x44\x63\x72\xfa\x40\x23\x61\x28\x9d\x38\x6f\x9b\x18\x3f\x24\x83\x3d\xbc\x1c\xf7\x18\x47\x59\x16\xbb\xc3\xdc\xbd\x8c\xef\xc4\x42\xc4\x10\xf1\x3a\x0f\xb7\x84\xad\xb9\x90\x33\xb5\x39\x63\x2d\x63\x9c\x35\x27\x05\x98\x56\x2e\xb2\xe6\x89\x80\xe1\xc5\xf4\x9c\x8d\x8f\x80\x2b\x50\xe3\xd9\x28\xbb\x89\x67\x45\xf4\xae\xc5\xdf\x57\x83\xbf\xf1\x6a\x26\x50\x58\xe6\x87\x76\xa1\xf9\x1b\x1b\x2a\xdc\x15\x2f\x8c\x6f\x1a\xe8\x65\xac\x91\xec\x0c\x58\x21\x95\x42\xcf\x42\xfb\x0e\x7f\x5e\x76\x56\x01\xb5\x15\xc7\x26\xba\x5c\x3b\xd7\xf6\x3d\x83\x77\xf7\xc1\x96\xc1\x3d\x9d\xeb\x6b\x18\x0a\xf5\x89\x5e\xd0\xd1\x7d\x98\xbb\x78\x89\x92\x9e\xb7\x59\x7e\x61\x1f\x6c\xce\x58\x7c\xa1\x51\xa8\x3b\xda\x1b\x28\x35\x58\x4e\x3c\x34\x62\xfa\xca\x7a\xff\xce\x9e\x6b\x64\x1f\x6c\x5d\x3c\xf2\x72\xf6\x05\x74\x2b\x70\x86\xcf\xcd\x1f\x19\xa5\xdf\x7d\x74\x6e\x62\x75\x58\xf2\xa4\x1d\xb2\x94\xd6\x00\x10\x9e\xa6\x84\xfd\x4d\x1c\x6e\x05\xb5\xc1\xf8\x9a\xf5\x38\xda\x5d\x81\xc3\x5f\x31\x25\xda\xee\x4a\xbb\x19\x84\xcc\x43\xc6\x68\xfe\xc8\xe1\xbe\xd1\xf5\xf6\x16\x21\x3c\xd2\x6b\x42\xce\xed\xe9\x31\x56\xc0\xff\xb3\xf4\x78\xc2\x2f\x81\x42\x0f\x4d\x1b\x9c\x61\xd2\x63\x74\xa8\x4b\x26\xb7\xad\xfa\xd4\x06\x97\x68\xbe\x91\x8f\x6e\xad\xb9\x9b\xd7\xe2\x3e\x5f\x7e\x89\xac\xf8\x77\xea\xff\x67\x4b\xfd\xd8\xa6\x93\x29\x91\xbd\xa6\x66\x49\x4f\xdd\x5c\x43\xc0\x44\xe9\x5e\xba\x74\xef\x52\xf2\x74\x2c\x32\x30\xb2\xa2\x1f\xe0\x54\xa6\xfc\xf7\x44\xbb\xa7\xe3\xf9\xf8\x69\x5b\xef\xd3\xde\x32\xa5\xbf\x26\xe8\xad\x5e\x3e\x7b\xd9\x12\xff\xde\x94\x2d\x4d\x7d\xf2\xbf\x01\x00\x00\xff\xff\x66\x13\x36\xf3\x95\x3f\x00\x00"),
          path: "jwt-api.tml",
          root: "jwt-api.tml",
        },
      
        "jwt-mock.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x5d\x6f\xeb\x36\x0c\x7d\xb6\x7e\x05\x17\x60\x80\x8d\x05\xbe\xd8\xf6\xd6\x21\x0f\x37\xb7\x6b\x51\x60\xe8\x8a\xad\x7b\x2a\x2e\x0a\x45\x62\x12\x23\x8e\x6c\x48\x4a\xd3\x20\xc8\x7f\x1f\x64\xca\xb1\xe3\x7c\xd8\x69\x52\xdc\x97\x16\xb1\xc4\x23\x1e\xf2\x90\xa2\xd8\x97\x2f\xf0\x20\x51\xd9\xc4\xae\x86\x5c\xcc\x50\x49\xd0\x68\x17\x5a\x19\xe0\xa0\x70\x09\x89\x32\x96\x2b\x81\x90\x8d\x61\xbd\x8e\x9f\xb8\x98\xf1\x09\x3e\xf2\x39\x6e\x36\x71\x69\x7a\x3b\xf4\xc6\x6c\xbc\x50\xa2\x89\x18\x46\x1d\x2c\x61\xcd\x82\x37\xae\x61\x9e\x89\x19\x6a\xa8\xd6\x19\x0b\xe4\x08\x6e\x06\x30\xe7\x33\x0c\xe7\x3c\x7f\x31\x56\x27\x6a\xf2\xfd\x28\x66\xc4\x58\x40\x30\xf1\xb7\x6c\xa1\xec\x9d\xf3\x69\x00\xce\xb5\x50\xd8\x77\x10\x99\xb2\xf8\x6e\xe3\x6f\xf4\x3f\x82\x30\x51\xb6\x0f\xa8\x75\xa6\x23\xe7\x47\x40\x21\x80\x14\x55\x28\x47\x51\x1f\x54\x92\xb2\x60\x53\x83\xd5\xc8\x2d\xb6\xe0\xf6\x01\x53\x9c\x1f\xa7\x1e\xd1\x89\xc5\x81\xc9\x18\x5e\xfb\x80\xef\x89\xb1\x8e\xaa\x1c\xbd\x38\xdb\xf8\x69\x31\x4a\x13\xf1\x70\xfb\xfd\x0f\xbf\xe6\xf6\x96\xde\x15\xd6\x26\x7e\xc4\x65\xd8\xd3\x28\x32\x2d\x81\xa7\x1a\xb9\x5c\xd1\x66\xd3\x8b\x58\x50\x78\x1d\xec\xe1\xc1\xa0\x70\xae\xa2\xda\x64\x78\x8b\x29\x76\x60\x98\x7b\x40\xa0\x9c\x9c\xa6\x94\x57\x6c\x7e\xda\xa7\xb3\x17\xa7\x3f\xb5\x7e\xcc\xec\x5d\xb6\x50\xb2\xe0\x11\x04\xb2\x70\x2a\x94\xa3\xea\xe0\xe8\x38\x85\x7b\xb4\x5f\xd3\x74\xb8\xfa\x5b\x4b\xd4\xad\x4c\x32\xb7\xcb\xd3\xf0\xbf\x86\xab\x2d\xad\xf0\xe5\xb8\xdc\x76\x94\xe3\x24\x4c\xc9\x30\x70\xc2\x86\x05\xc1\x38\xd3\x45\x80\x9c\x46\x6e\x06\xa0\xb9\x9a\x20\xc8\x51\x19\x13\x82\x18\x00\xcf\x73\x57\x42\xfe\x03\xed\x8f\x7c\x3c\x3c\xf1\xed\xda\xe1\x08\x5c\x44\xbd\x0f\x39\x9f\x20\x14\x05\xa2\xd1\xe4\x99\x32\xf8\x84\xfa\xc9\x7f\x6c\x8b\x4c\xb3\xb0\x7e\x68\x78\x5c\x39\xfb\x1f\x07\x6a\xfa\x1e\xdb\x1a\xc5\x01\xb9\x87\xdd\x54\xe1\xbc\x3a\x52\x0b\x54\x29\x07\xea\x81\x6c\x3a\x54\xc5\xce\xfe\x03\xac\x86\xab\xbb\x04\x53\xd9\x4a\x6e\x86\x55\xd2\xdf\x78\xba\x28\x12\x8c\x7a\xcc\x05\xae\x37\x9d\xa9\xbe\xf1\xd4\x51\x1c\xcf\x6d\xfc\x6f\xae\x13\x65\xc7\x61\xef\xe7\x5f\x4c\xcf\x63\x46\x55\x6a\x7d\xcf\xda\x4b\xae\x59\x26\x56\x4c\xbd\x2b\x26\x7e\xce\xfe\xca\x96\xa8\xc3\x19\xae\xe8\x84\x40\x70\x83\xd0\xd3\x38\xd6\x68\xa6\xaf\x36\x9b\xa1\xea\xdd\xb8\x05\x17\x49\x42\x8d\xff\xa1\xd5\x67\xb7\x08\x83\x81\x3b\x9c\x8c\x1b\xb2\xf0\x11\x0b\x28\x94\x1e\xda\x72\x3d\x41\xfb\x9a\xc8\x3d\xd8\xe7\x62\xe5\xe1\xf6\x6c\x48\xca\xf7\x21\xc8\xb2\x27\x9f\x01\xb9\x29\xfb\xfa\xb1\xd6\x59\x66\x66\xbd\x69\x53\x50\x4d\x2a\xff\xe5\xb2\xcb\xa5\xd6\xa8\x81\xcb\x6e\xb9\x0b\xaf\x04\xba\xda\xf2\xf6\x5b\xcd\xff\x26\xa6\x6c\xc3\xea\x73\xcf\xed\x10\x24\x8e\x13\x85\x6e\xe4\x11\x99\x12\x1a\x2d\x3a\x76\x0b\x61\x61\x39\x4d\xc4\x14\x92\x79\x9e\xe2\x1c\x95\x35\x60\xa7\x08\x73\xb4\xd3\x4c\x1a\x70\x42\xb6\x53\x74\x60\x1d\x06\x9c\x6d\x35\xc5\xf0\x35\x4d\xb7\x20\xcb\x24\x4d\x21\xe7\x2a\x11\x7d\x30\x19\x70\x29\x41\xa1\x40\x63\xb8\x5e\x91\x8d\xe2\x29\xa4\xd9\x24\x11\x31\xb3\xab\x1c\xeb\x8e\x7b\x2f\xd7\x2c\xa8\x26\x9d\x6e\x73\x0e\x63\x41\xed\x8e\x3f\xf7\x86\x67\x2c\xa8\xcd\x40\x17\x4e\x40\x8c\x05\x65\xf3\xbd\x6a\xeb\x65\x2c\xa8\x69\xfa\xda\x8a\x26\xaf\x1b\x13\xc6\x67\xce\x17\x74\x60\xbd\x9b\x7f\x62\x2f\xdf\xb2\xbb\x80\xd6\xf5\x66\x07\x5f\xb1\x85\xc6\x9b\xc5\xb8\x57\x68\x34\xf4\x87\x91\xaf\xb0\xa2\x4a\xab\x4d\x31\x3d\x51\x42\x87\x52\xfb\x1c\x11\x78\xa7\xf7\x41\xd1\x4e\x7e\x75\x9c\xec\x6f\xae\x8f\x39\xa8\xea\xa5\xe1\x20\x22\x56\xf5\xef\x72\x27\xf3\x24\xa8\xea\xda\x59\xd0\xbe\xb3\x69\x78\xb3\xf3\xa6\x75\x62\xb4\xe5\x52\x35\x06\x87\xb3\x33\x6c\xd7\x68\x95\x84\xa8\x11\x74\x48\x4b\xb1\xef\xfc\xbc\x90\xd9\x45\x0f\xac\x06\xbf\xaa\x75\x11\x3f\x3f\x34\x1e\xe0\x76\x8f\x1d\xf4\x76\x8f\xe7\xab\xcd\xd9\x5c\x71\xc4\x3c\x28\x49\xdf\x54\x4f\xe5\x70\x57\x9a\xd4\x2d\xdb\x09\xd3\xbe\xb3\x39\x7b\xb3\xeb\x4e\x15\x8d\xd4\x56\x0d\x7f\x97\x76\x4b\x92\xab\x36\xde\x29\xdd\xd5\xf6\x8f\x24\xbe\x66\xfd\x79\x4f\xd1\x63\x8a\xd8\xbd\xb0\x28\x48\xc5\x11\xdb\x93\x4e\x29\xa4\xba\x80\x3a\xc5\xc9\xef\xfd\x48\x90\x4a\xd3\x4f\x7b\xaa\x1c\x0b\x50\xed\x82\xa5\xe8\xcc\x70\x55\x3d\x5e\x4e\x45\xc6\x4d\x75\x1d\xd5\xf3\x41\xd9\xfc\xe0\xf7\x7b\x3d\x64\xc5\xdf\xdf\x1b\xca\x3a\x26\x29\xf2\x62\xcf\x83\x43\xf1\x24\xdc\x22\xaa\xff\x07\x00\x00\xff\xff\xa7\x2f\x9e\xd1\x9d\x14\x00\x00"),
          path: "jwt-mock.tml",
          root: "jwt-mock.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
