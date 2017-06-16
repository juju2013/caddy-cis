# caddy-cis #

Make Path Case InSensitive, middleware for [Caddy Server](https://github.com/mholt/caddy)

## Use Case

To serve `http://yourserver/path/to/file.txt` , you have `${DOCUMENTROOT}/PATH/to/FILE.txt`, 
and change file name is not an option because:

  * You've billions of them
  * Changes are replicated over sea and take months
  * Etc...

## Plugin Usage

```
 :80 {
   root ${DOCUMENTROOT}
   caseinsensitive
   ...
}
```

For `${DOCUMENTROOT}/AAA/BBB/CCC/.../FILE.EXT`:

  * `AAA`,`BBB`,`CCC`,...,`FILE` and `EXT` must be all upper or lower, but not necessary the same
  * Request URL can be in any case

Works for:

```
  ${DOCUMENTROOT}/AAA/bbb/CCC/.../FILE.ext
  ${DOCUMENTROOT}/aaa/BBB/ccc/.../file.EXT
  ${DOCUMENTROOT}/aaa/bbb/ccc/.../file.ext
```


Does not work for:

```
  ${DOCUMENTROOT}/AaA/bBB/CCc/.../FiLE.ext
```


### Known limitations ###
To find the right case/path, os.Stat() is used against each folder, filename and extension. 
In the worest case, there will have (3 x number of folders + 4) calls to os.Stat for each URL.

All errors - permission problem for example - will be treated as file does not exist.

### Acknowledgement ###
Courtesy from [Filae](https://www.filae.com)
