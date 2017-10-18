import json
import requests

import config

def create_gust(code):
    url = config.WEB_URL + '/auth/guest'
    data = {
        "Mobilecode": code
    }

    r = requests.post(url, json=data, headers={'Bundle-Identifier':config.GAME["NAME"]})
    return  json.loads(r.text)