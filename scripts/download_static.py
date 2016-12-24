import os
import requests
import zipfile
import shutil

bootstrap_downl_link = "https://github.com/twbs/bootstrap/releases/download/v4.0.0-alpha.5/bootstrap-4.0.0-alpha.5-dist.zip"
fa_downl_link = "http://fontawesome.io/assets/font-awesome-4.7.0.zip"
fantasque_sans_downl_link = "https://fontlibrary.org/assets/downloads/fantasque-sans-mono/b0cbb25e73a9f8354e96d89524f613e7/fantasque-sans-mono.zip"

def download_static_root(static_path):
    final_bootstrap_path = os.path.join(static_path, "css", "bootstrap.min.css")
    final_fontawesome_path = os.path.join(static_path, "css", "font-awesome-4.7.0")
    final_fantasque_path = os.path.join(static_path, "fonts", "fantasque-sans-mono")

    css_path = create_css_dir(static_path)
    fonts_path = create_fonts_dir(static_path)

    get_bootstrap(css_path, final_bootstrap_path)
    get_font_awesome(css_path, final_fontawesome_path)
    get_fantasque_sans(fonts_path, final_fantasque_path)

def create_css_dir(static_path):
    css_path = os.path.join(static_path, "css")
    if not os.path.isdir(css_path):
        os.makedirs(css_path)
    return css_path

def create_fonts_dir(static_path):
    fonts_path = os.path.join(static_path, "fonts")
    if not os.path.isdir(fonts_path):
        os.makedirs(fonts_path)
    return fonts_path

def get_bootstrap(css_path, final_path):
    if os.path.exists(final_path):
        return
    bootstrap_downl = os.path.join(css_path, "bootstrap.zip")

    r = requests.get(bootstrap_downl_link, allow_redirects=True)
    if not r.ok:
        print("get_bootstrap() failed")
    else:
        with open(bootstrap_downl, 'wb') as fd:
            for chunk in r.iter_content(chunk_size=128):
                fd.write(chunk)

    zip_ref = zipfile.ZipFile(bootstrap_downl, 'r')
    zip_ref.extractall(css_path)
    zip_ref.close()
    bootstrap_extract = ""
    for root, dirs, files in os.walk(css_path):
        for dir in dirs:
            if dir.startswith("bootstrap"):
                bootstrap_extract = dir

    min_css_path = os.path.join(css_path, bootstrap_extract, "css", "bootstrap.min.css")
    shutil.copy(min_css_path, final_path)

    os.remove(bootstrap_downl)
    shutil.rmtree(os.path.join(css_path, bootstrap_extract))

def get_font_awesome(css_path, final_path):
    if os.path.exists(final_path):
        return
    fontawesome_downl = os.path.join(css_path, "fontawesome.zip")

    r = requests.get(fa_downl_link, allow_redirects=True)
    if not r.ok:
        print("get_font_awesome() failed")
    else:
        with open(fontawesome_downl, 'wb') as fd:
            for chunk in r.iter_content(chunk_size=128):
                fd.write(chunk)

        zip_ref = zipfile.ZipFile(fontawesome_downl, 'r')
        zip_ref.extractall(css_path)
        zip_ref.close()
        os.remove(fontawesome_downl)

def get_fantasque_sans(fonts_path, final_path):
    if os.path.exists(final_path):
        return
    fantasque_downl = os.path.join(fonts_path, "fantasque.zip")

    r = requests.get(fantasque_sans_downl_link, allow_redirects=True)
    if not r.ok:
        print("get_fantasque_sans() failed")
    else:
        with open(fantasque_downl, 'wb') as fd:
            for chunk in r.iter_content(chunk_size=128):
                fd.write(chunk)

        zip_ref = zipfile.ZipFile(fantasque_downl, 'r')
        zip_ref.extractall(final_path)
        zip_ref.close()
        os.remove(fantasque_downl)
