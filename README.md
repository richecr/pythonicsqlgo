# Pythonic Core (Under development)


## Go to Python3

#### Generate package Python

```sh
gopy pkg -name=pythonic_core -rename=true -author="Rich Ramalho" -email="richelton14@gmail.com" -desc="The core of pythonic in Golang" -url="https://github.com/richecr/pythonic-core" -output=pythonic_core -vm=python3 github.com/richecr/pythonic-core github.com/richecr/pythonic-core/lib/dialects github.com/richecr/pythonic-core/lib/query
```

#### Generate dist

```sh
python setup.py bdist_wheel
```

To publish to PyPi you need to use this command:

```sh
python -m build
```

#### Install local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip install $wheel_file
```

#### Uninstall local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip uninstall $wheel_file
```
