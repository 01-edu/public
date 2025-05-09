import time
from datetime import datetime
from random import uniform

from flask import Flask, jsonify, request
from flask_cors import CORS

from utils import get_historical_data, load_data

app = Flask(__name__)
CORS(app)

start_time = time.time()

historical_data = load_data()


@app.route('/exchange_rate/<symbol>')
def get_stock_data(symbol):
    if symbol not in list(historical_data.keys()):
        response = jsonify({"error": "Invalid symbol"})
        response.status_code = 404
        return response
    current_time = time.time()
    last_value = historical_data[symbol].iloc[-1].Close
    step = (int(current_time * 10) - int(start_time * 10)
            ) % len(historical_data[symbol])
    return jsonify({
        'currency': 'USD',
        'rate': last_value * (1 + uniform(0.05, -0.05) + step * 0.0005),
        'datetime': datetime.fromtimestamp(current_time)
    })


@app.route('/hist/<symbol>')
def get_hist_data(symbol):
    if symbol not in list(historical_data.keys()):
        response = jsonify({"error": "Invalid symbol"})
        response.status_code = 404
        return response

    df = historical_data[symbol]
    args = request.args
    start_date = args.get("start_date")
    end_date = args.get("end_date")
    if not start_date or not end_date:
        response = jsonify({"error": "start_date and end_date required"})
        response.status_code = 400
        return response

    try:
        filtered = get_historical_data(df, start_date, end_date, start_time)
        data = filtered[['datetime', 'Close']].to_dict(orient="list")

        values = [
            {"date": dt.date(), "close": close}
            for dt, close in zip(data["datetime"], data["Close"])]
        response = jsonify({
            "currency": 'USD',
            "values": values,
        })
        return response

    except Exception as e:
        response = jsonify({'error': str(e)})
        response.status_code = 400
        return response


@app.route('/stocks_list')
def list_symbols():
    return jsonify(list(historical_data.keys()))


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
