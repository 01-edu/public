from flask import Flask, jsonify
from flask_cors import CORS

import time
from utils import load_historical_data

app = Flask(__name__)
CORS(app)

start_time = time.time()

historical_data = load_historical_data()


@app.route('/exchange_rate/<symbol>')
def get_stock_data(symbol):
    if symbol not in list(historical_data.keys()):
        return jsonify("Invalid symbol")
    current_time = time.time()
    step = (int(current_time * 10) - int(start_time * 10)
            ) % len(historical_data[symbol])
    try:
        return jsonify({
            'symbol': 'USD',
            'rate': float(historical_data[symbol][step]),
            'timestamp': current_time
        })
    except:
        return "Server Error"


@app.route('/stocks_list')
def list_symbols():
    return jsonify(list(historical_data.keys()))


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
