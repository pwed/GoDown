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
  downloadChecksum: string;
  downloadHashType: string = "none";

  hashTypes = [
    {value: 'none', viewValue: 'None'},
    {value: 'md5', viewValue: 'MD5'},
    {value: 'rsa', viewValue: 'RSA'},
    {value: 'sha1', viewValue: 'SHA1'},
    {value: 'sha256', viewValue: 'SHA256'},
    {value: 'sha512', viewValue: 'SHA512'},
  ];

  readonly ROOT_URL = "/api/"

  response: any;


  ngOnInit() {
  }

  startDownload() {
    console.log(this.downloadURL)
    this.http.post(this.ROOT_URL + 'startDownload',
    JSON.stringify({downloadURL: this.downloadURL, downloadChecksum: this.downloadChecksum, hashType: this.downloadHashType})).subscribe()
  }

}
