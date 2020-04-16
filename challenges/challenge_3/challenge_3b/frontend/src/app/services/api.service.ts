import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  constructor(private http: HttpClient) {}
  // private URL = 'http://localhost:9090/webapi/v1/user';
  private URL = 'https://jsonplaceholder.typicode.com/posts/';
  public result;

  getAll() {
    return this.http.get(this.URL);
  }
}
