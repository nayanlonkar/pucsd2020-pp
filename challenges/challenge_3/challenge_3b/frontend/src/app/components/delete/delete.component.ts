import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-delete',
  templateUrl: './delete.component.html',
  styleUrls: ['./delete.component.css'],
})
export class DeleteComponent implements OnInit {
  constructor(private http: HttpClient) {}

  // private URL = 'https://jsonplaceholder.typicode.com/posts/';
  // private URL = 'http://localhost:9090/webapi/v1/user';
  private URL = '/api/webapi/v1/user/';

  public result;

  ngOnInit(): void {}

  async onClickDelete(value) {
    console.log(value);
    this.result = await this.http.delete(this.URL + value).toPromise();
  }
}
