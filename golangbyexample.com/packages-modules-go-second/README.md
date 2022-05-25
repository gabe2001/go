1. mkdir math
2. cd math
3. create math.go
4. go mod init my/matte
   1. you can name your module any which way you like
5. cd ..
6. mkdir school
7. cd school
8. go mod init school
9. edit go.mod and add ```replace my/matte => ../math``` at end of file
      1. ```replace``` points to a __directory__ not a __module__ nor a __package__
10. when you try running school.go you should get an error message - it hints to next step
11. ```go get my/matte```
    1. it will add ```require my/matte v0.0.0-00010101000000-000000000000 // indirect``` to go.mod
12. create school.go
    1. notice that the import requires an alias to map the __directory__ to the __module__