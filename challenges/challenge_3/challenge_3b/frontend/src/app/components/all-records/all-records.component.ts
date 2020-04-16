import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-all-records',
  templateUrl: './all-records.component.html',
  styleUrls: ['./all-records.component.css'],
})
export class AllRecordsComponent implements OnInit {
  constructor(private apiservice: ApiService) {}
  public result;
  ngOnInit(): void {
    // this.apiservice.getAll().subscribe((data) => console.log(data));
    this.apiservice.getAll().subscribe((data) => (this.result = data));
  }
}
