from flask import Flask, jsonify, request
from flask_cors import CORS

from random import uniform
import time
from utils import load_data

app = Flask(__name__)
CORS(app)

start_time = time.time()

historical_data = load_data()


@app.route('/exchange_rate/<symbol>')
def get_stock_data(symbol):
    if symbol not in list(historical_data.keys()):
        return jsonify("Invalid symbol")
    current_time = time.time()
    last_value = historical_data[symbol].Close[-1]
    step = (int(current_time * 10) - int(start_time * 10)
            ) % len(historical_data[symbol])
    return jsonify({
        'symbol': 'USD',
        'rate': last_value * (1 + uniform(0.05, -0.05) + step * 0.0005),
        'timestamp': current_time
    })


@app.route('/stocks_list')
def list_symbols():
    return jsonify(list(historical_data.keys()))


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
