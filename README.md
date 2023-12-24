# PythonicSQL (Under development)


## Go to Python3

#### Generate package Python

```sh
gopy pkg -name=pythonicsql -rename=true -author="Rich Ramalho" -email="richelton14@gmail.com" -desc="The unofficial HLTV Python API" -url="https://github.com/richecr/pythonicsqlgo" -output=pythonicsql -vm=python3 github.com/richecr/pythonicsqlgo github.com/richecr/pythonicsqlgo/lib/pythonic github.com/richecr/pythonicsqlgo/lib/query github.com/richecr/pythonicsqlgo/lib/dialects
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
