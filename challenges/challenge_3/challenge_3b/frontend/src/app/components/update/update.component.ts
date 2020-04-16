import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css'],
})
export class UpdateComponent implements OnInit {
  constructor(private http: HttpClient) {}

  // private URL = 'https://jsonplaceholder.typicode.com/posts/';
  // private URL = 'http://localhost:9090/webapi/v1/user';
  private URL = '/api/webapi/v1/user';
  public result;
  public data = {
    first_name: null,
    last_name: null,
    email: null,
    contact_number: null,
    password: null,
    updated_by: 0,
  };

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*',
    }),
  };

  ngOnInit(): void {}

  async onSubmitUpdate(values) {
    this.data.first_name = values.first_name;
    this.data.last_name = values.last_name;
    this.data.email = values.email_id;
    this.data.contact_number = values.contact_number;
    this.data.password = values.password;

    this.result = await this.http
      .put(
        this.URL + '/' + values.id,
        JSON.stringify(this.data),
        this.httpOptions
      )
      .toPromise();
  }
}
