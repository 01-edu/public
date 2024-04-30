import os
import csv
import pandas as pd


def load_historical_data(directory_path='./sample-stocks/'):
    historical_data = {}

    file_list = [filename for filename in os.listdir(
        directory_path) if filename.endswith(".csv")]
    for filename in file_list:
        symbol = filename.replace(".csv", "")

        historical_data[symbol] = {}
        file_path = os.path.join(directory_path, filename)
        with open(file_path, 'r') as csv_file:
            csv_reader = csv.DictReader(csv_file)
            historical_data[symbol] = [row['Close'] for row in csv_reader]

    return historical_data


def load_data(directory_path='./sample-stocks/'):
    historical_data = {}

    file_list = [filename for filename in os.listdir(
        directory_path) if filename.endswith(".csv")]
    for filename in file_list:
        symbol = filename.replace(".csv", "")
        file_path = os.path.join(directory_path, filename)
        historical_data[symbol] = pd.read_csv(
            file_path, index_col=0, parse_dates=True)

    return historical_data


if __name__ == "__main__":
    result = load_data()
    print(f'keys: {result.keys()}')
    print(result["AE"].Close[-1])
