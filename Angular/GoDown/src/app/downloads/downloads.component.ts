import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-downloads',
  templateUrl: './downloads.component.html',
  styleUrls: ['./downloads.component.scss']
})
export class DownloadsComponent implements OnInit {

  constructor(private http: HttpClient) { }

  downloadURL: string;

  readonly ROOT_URL = "/api/"

  response: any;


  ngOnInit() {
  }

  startDownload() {
    console.log(this.downloadURL)
    this.http.post(this.ROOT_URL + 'startDownload', this.downloadURL).subscribe()
  }

}
