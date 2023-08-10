# Use the ubuntu
FROM python:3.10-slim

# Set the working directory to /app
WORKDIR /app

COPY sample-stocks/ ./sample-stocks/

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt 

COPY ["app.py", "utils.py",  "./"] 

CMD ["python3", "app.py"]
