from bs4 import BeautifulSoup
import requests
import json


def main(event, context):
    '''
    Scrape the page and returns the titles of the items which are displayed for sale.
    '''
    url = event['data']['url']
    r = requests.get(url)
    soup = BeautifulSoup(r.text, 'html.parser')

    # https://webscraper.io/test-sites/e-commerce/allinone
    items =  [
        x.find("p",{"class":"description"}).text 
        for x in soup.find_all("div",{"class":["col-sm-4", "col-lg-4", "col-md-4"]})
    ]

    return json.dumps(items)