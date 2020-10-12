FUCTION_PACKAGE_FILE = function_package.zip


function_package:
	# Install scrapper dependencies locally
	cd scrapper && pip3 install -t . bs4 requests 

	# Package them into zip folder
	cd scrapper && zip -r9 ./${FUCTION_PACKAGE_FILE} *

	# Move the package to upper folder
	mv scrapper/${FUCTION_PACKAGE_FILE} ./${FUCTION_PACKAGE_FILE}