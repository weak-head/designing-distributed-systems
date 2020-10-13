SCRAPER_PACKAGE_FILE = scraper_package.zip

.PHONY: scraper_package
scraper_package:
	# Install scraper dependencies locally
	cd scraper && pip3 install -t . bs4 requests 

	# Package function with dependencies
	cd scraper && zip -r9 ./${SCRAPER_PACKAGE_FILE} *

.PHONY: clean
clean:
	# Delete scraper function
	kubeless function delete scraper

.PHONY: deploy
deploy:
	# Deploy scrapper function
	kubectl apply -f scraper.yaml

	# kubectl get functions
	# kubeless function ls
	# kubeless function call scraper --data '{"url":"https://webscraper.io/test-sites/e-commerce/allinone"}'