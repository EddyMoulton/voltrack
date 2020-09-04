import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { OwnedStockViewModel } from '@/models/OwnedStockViewModel';
import { AddTransactionDto } from './models/AddTransactionDto';
import { TransactionSummaryViewModel } from '@/models/TransactionSummaryViewModel';
import { TransactionSummaryDto } from '@/models/TransactionSummaryDto';
import { StockLogViewModel } from '@/models/StockLogViewModel';

export class ApiClient {
  private readonly getOwnedStockLogsUrl = '/stocks/history';
  private readonly generateSummaryLogsUrl = '/reporting/generate';
  private readonly getAllTransactionsUrl = '/stocks/transactions';
  private readonly addTransactionUrl = '/stocks/transactions';
  private readonly addTransactionsUrl = '/stocks/transactions/bulk';
  private readonly getAllStocksUrl = '/stocks';
  private readonly getCurrentStocksUrl = '/stocks/current';
  private readonly addStockLogsUrl = '/stocks/logs';
  private readonly getTransactionSummariesUrl = '/transactions/summaries';

  private axios: AxiosInstance;
  private baseUrl: string;

  constructor(baseUrl?: string, axiosConfig?: AxiosRequestConfig | undefined) {
    this.baseUrl = baseUrl || process.env.VUE_APP_API_URL || '';
    this.baseUrl = this.baseUrl.replace(/\/$/, ''); // Ensure no trailing slash

    if (!axiosConfig) {
      axiosConfig = {};
    }

    axiosConfig.headers = {
      ...axiosConfig?.headers,
      ...{
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Headers':
          'access-control-allow-origin, access-control-allow-headers',
        'Content-Type': 'text/plain',
      },
    };

    this.axios = axios.create(axiosConfig);
  }

  public getOwnedStockLogs() {
    return this.axios.get(this.baseUrl + this.getOwnedStockLogsUrl);
  }

  public GenerateSummaryLogs() {
    return this.axios.get(this.baseUrl + this.generateSummaryLogsUrl);
  }

  public GetAllTransactions() {
    return this.axios.get(this.baseUrl + this.getAllTransactionsUrl);
  }

  public AddTransaction(dto: AddTransactionDto) {
    return this.axios.post(this.baseUrl + this.addTransactionUrl, dto);
  }

  public AddTransactions(data: AddTransactionDto[]) {
    data.forEach((transaction) => {
      if (transaction.date) {
        const val = new Date(transaction.date);
        transaction.date = val;
      }

      if (transaction.fee) {
        transaction.fee = Math.floor(transaction.fee * 10000);
      }

      if (transaction.cost) {
        transaction.cost = Math.floor(transaction.cost * 10000);
      }
    });

    return this.axios.post(this.baseUrl + this.addTransactionsUrl, {
      transactions: data,
    });
  }

  public GetAllStocks() {
    return this.axios.get(this.baseUrl + this.getAllStocksUrl);
  }

  public GetCurrentStocks(): Promise<OwnedStockViewModel[]> {
    return this.axios
      .get(this.baseUrl + this.getCurrentStocksUrl)
      .then((result) => {
        if (result.status == 200) {
          return result.data.currentStocks.map((data: OwnedStockViewModel) => {
            const entry = data as OwnedStockViewModel;

            entry.currentValue = entry.currentValue / 10000;
            entry.paidValue = entry.paidValue / 10000;
            entry.difference = entry.difference / 10000;
            entry.totalDividends = entry.totalDividends / 10000;

            return entry;
          });
        } else {
          throw Error(result.statusText);
        }
      });
  }

  public async GetTransactionSummaries(): Promise<
    TransactionSummaryViewModel[]
  > {
    const result = await this.axios.get(
      this.baseUrl + this.getTransactionSummariesUrl,
    );

    const totalCost = result.data.transactions
      .map((t: TransactionSummaryDto) => t.cost)
      .reduce((total: number, cost: number) => {
        return total + cost || 0;
      });

    const data = result.data.transactions.map((t: TransactionSummaryDto) => {
      const viewModel = new TransactionSummaryViewModel();
      viewModel.generate(
        t.code || '',
        t.quantity || 0,
        ((t.cost || 0) * (t.quantity || 0)) / 10000,
        ((t.value || 0) * (t.quantity || 0)) / 10000,
        (t.dividendValue || 0) / 10000,
        new Date(t.date || 0),
        totalCost / 10000,
      );
      return viewModel;
    });

    return data;
  }

  public async AddStockLogs(logs: StockLogViewModel[]) {
    logs.forEach((log) => {
      if (log.date) {
        const val = new Date(log.date);
        log.date = val;
      }

      if (log.value) {
        log.value = Math.floor(log.value * 10000);
      }
    });

    return await this.axios.post(this.baseUrl + this.addStockLogsUrl, { logs });
  }
}
