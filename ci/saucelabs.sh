#!/bin/bash -e

echo "Installing Pip"
cd `mktemp -d`
wget https://bootstrap.pypa.io/get-pip.py
python get-pip.py
cd $SD_SOURCE_DIR

echo "Installing Selenium and SauceClient"
pip install selenium sauceclient

echo "Running SauceLab tests"
python ui-tests.py
