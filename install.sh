gopy pkg -name=pythonic_core -rename=true -author="Rich Ramalho" -email="richelton14@gmail.com" -desc="The core of pythonic in Golang" -url="https://github.com/richecr/pythonic-core" -output=pythonic_core -vm=python3 github.com/richecr/pythonic-core github.com/richecr/pythonic-core/lib/dialects github.com/richecr/pythonic-core/lib/query

cd pythonic_core/
python setup.py bdist_wheel
python -m build

wheel_file=$(ls dist/*.whl | head -n1); pip install $wheel_file