import os
import json
import requests
from requests.auth import HTTPBasicAuth

def index_json_to_server(json_folder_path, api_url, username, password):
    for json_file in os.listdir(json_folder_path):
        if json_file.endswith(".json"):
            json_file_path = os.path.join(json_folder_path, json_file)
            with open(json_file_path, 'r', encoding='utf-8') as f:
                json_content = f.read()
                
            response = requests.post(
                api_url,
                auth=HTTPBasicAuth(username, password),
                headers={'Content-Type': 'application/json'},
                data=json_content,
                verify=False  # Deshabilitar la verificación SSL, ya que se está utilizando -k en curl
            )
            if response.status_code == 200:
                print(f"Successfully indexed {json_file}")
            else:
                print(f"Failed to index {json_file}: {response.status_code}, {response.text}")

# Parámetros de la API
api_url = "https://api.openobserve.ai/api/terlyn_organization_418/default/_json"
username = "galeanoterlyn@gmail.com"
password = "348n2GsAw5IU96EK170o"

# Ruta de la carpeta que contiene los archivos JSON
json_folder_path = '/Users/macbook/Documents/proyecto_v2_mail/Data/bachs'

# Indexar los archivos JSON al servidor
index_json_to_server(json_folder_path, api_url, username, password)
