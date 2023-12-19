@echo off

python -m venv venv
call ./venv/Scripts/activate
pip install -r requirements.txt
cls
python ./app/main.py

cmd /k