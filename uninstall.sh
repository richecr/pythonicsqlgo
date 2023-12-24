cd pythonicsql
wheel_file=$(ls dist/*.whl | head -n1); pip uninstall $wheel_file
cd ..
rm -rf pythonicsql/