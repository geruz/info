go build ./simple.go
echo -e '\nversion output: '
./simple  --version
echo -e '\napp version and lib output:'
./simple --version-all
echo -e '\napp instance id:'
./simple --version-instance
echo -e '\n'
