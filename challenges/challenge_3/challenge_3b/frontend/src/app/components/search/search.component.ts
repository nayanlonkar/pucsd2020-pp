import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css'],
})
export class SearchComponent implements OnInit {
  constructor(private http: HttpClient) {}
  public result = null;
  // private URL = 'https://jsonplaceholder.typicode.com/posts/';
  // private URL = 'http://localhost:9090/webapi/v1/user';
  private URL = '/api/webapi/v1/user/';

  ngOnInit(): void {}

  async onClickSubmit(value) {
    this.result = await this.http.get(this.URL + value.id).toPromise();
    this.result = this.result.data;
    console.log(this.result);
  }
}
