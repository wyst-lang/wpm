import urllib
import requests
import os
import shutil
import zipfile
from sys import argv
from urllib.parse import urlparse

def extract(zip_file, pkgName):
    os.makedirs("tmp", exist_ok=True)
    with zipfile.ZipFile(zip_file, 'r') as zip_ref:
        zip_ref.extractall(path="tmp")
    old_folder_name = os.listdir("tmp")[0]
    final_destination = 'lib/'+pkgName
    try:shutil.rmtree(final_destination)
    except:pass
    shutil.move(os.path.join("tmp", old_folder_name), final_destination)
    shutil.rmtree('tmp')

def downloadRelease(pkgName, repo, tag):
    if tag=='latest':
        tag = requests.get(os.path.join('https://api.github.com/repos', 
            str(urlparse(repo).path).removeprefix('/'),
            'releases/latest'),
        headers={"Accept": "application/vnd.github+json"}).json()['tag_name']
    url = os.path.join(repo, f'archive/refs/tags/{tag}.zip')
    urllib.request.urlretrieve(url, "temp.zip")
    input_zip_path = 'temp.zip'
    new_folder_name = pkgName
    extract(input_zip_path, new_folder_name)
    os.remove('temp.zip')

def installPackage(packageName, packageVersion='latest'):
    print(f"Indexing package {packageName}")
    req = requests.get("http://localhost:3000", json={'name': packageName})
    data = req.json()
    if req.status_code==200:
        print(f"Downloading package {packageName}")
        downloadRelease(packageName, data['repo'], packageVersion)
        print(f"Done")
    else:
        print(f'Error ({packageName}): {data["message"]}')
        quit(1)

if __name__=="__main__":
    argl = len(argv)
    if argl > 1:
        match argv[1]:
            case "install":
                if argl > 2:
                    for pkg in argv[2::]:
                        package = pkg.split('@')
                        if len(package) == 2:
                            installPackage(package[0], package[1])
                        else:
                            installPackage(pkg)
            case _:...