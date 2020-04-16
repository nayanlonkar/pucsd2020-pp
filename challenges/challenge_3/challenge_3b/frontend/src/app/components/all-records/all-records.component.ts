import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-all-records',
  templateUrl: './all-records.component.html',
  styleUrls: ['./all-records.component.css'],
})
export class AllRecordsComponent implements OnInit {
  constructor(private http: HttpClient) {}

  // private URL = 'https://jsonplaceholder.typicode.com/posts/';
  // private URL = 'http://localhost:9090/webapi/v1/user';
  private URL = '/api/webapi/v1/user';

  public result;
  async ngOnInit() {
    this.result = await this.http.get(this.URL).toPromise();
    this.result = this.result.data;
    console.log(this.result[0]);
  }
}
