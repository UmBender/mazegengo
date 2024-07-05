# Mazegengo
A maze generator made with go, using raylib to generate pretty images.

## Algorithm 
The algorithm to generate the maze is simple. It only uses a random dfs.
The algorithm starts in some point, in this case it only starts in the superior left side,
and tries to acess all four neighbours in a random order, it only acess that cell if it wasn't acessed before.


### Build
You must have installed the go compiler.
Than run the folowing comand, if you're using linux.
```bash
go build main.go
```
