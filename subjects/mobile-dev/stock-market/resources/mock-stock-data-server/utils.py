import os
from datetime import datetime

import pandas as pd


def load_data(directory_path='./sample-stocks/'):
    historical_data = {}

    file_list = [filename for filename in os.listdir(
        directory_path) if filename.endswith(".csv")]
    for filename in file_list:
        symbol = filename.replace(".csv", "")
        file_path = os.path.join(directory_path, filename)
        historical_data[symbol] = pd.read_csv(file_path, parse_dates=[0])
    return historical_data


def get_historical_data(df, start, end, start_time):
    today = datetime.fromtimestamp(start_time).date()
    last_entry = df.sort_values(by="Date").iloc[-1]
    delta_today_last_entry = today - last_entry.Date.date()

    try:
        query_start_dt = datetime.fromisoformat(start)
        query_end_dt = datetime.fromisoformat(end)
        if query_end_dt < query_start_dt:
            raise Exception("end_date must come after start_date")
        if query_end_dt.date() > today:
            query_end_dt = datetime.fromtimestamp(start_time)

        df['datetime'] = df.Date + delta_today_last_entry
        return (df.loc[
              (df.datetime >= query_start_dt) &
              (df.datetime <= query_end_dt)])
    except Exception as e:
        raise Exception(str(e))


if __name__ == "__main__":

    result = load_data()
    print(f'keys: {result.keys()}')
    now = datetime.now()

    df = result["AE"]
    print(df.info())
