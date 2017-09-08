# zip_dict_cracker
Dictionary based zip file password cracker in go

### Usage
```
Usage: zip_dict_cracker [options]

  -d string
        Path to dictionay file
  -f string
        Path to zip file
```
Example:
```
.\zip_dict_cracker -f test.zip -d passwords.txt
```
Output:
```
Checked 100000 passwords. Rate: 17243.61 passwords per second
Password found: 1234512345
```

License:
--------
Released under [The MIT License](https://github.com/delimitry/zip_dict_cracker/blob/master/LICENSE).
