import requests

def main():
    for i in range(10000):
        requests.post("http://127.0.0.1:8000/note",{'content': '呵呵'+str(i)})

main()