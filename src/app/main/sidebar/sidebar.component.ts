import { Component, OnInit, ViewChild } from '@angular/core';
import { MainService } from '../../services/main.service';
import { TestResult } from '../../models/testresult';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {
  results: TestResult;

  constructor(public mainService: MainService) { }

  ngOnInit() {
  }

  run() {
    this.mainService.runTests(`TestReverseString`, this.mainService.solution)
      .subscribe(res => { this.results = res['body']; }, err => console.log(err));
  }

  format() {
    this.mainService.fmt(this.mainService.solution);
  }
}
