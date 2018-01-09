import { Component, OnInit } from '@angular/core';
import { MainService } from '../../services/main.service';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.css']
})
export class EditorComponent implements OnInit {
  initialcode: string;
  solution: string;

  constructor(public mainService: MainService) { }

  ngOnInit() {
  }

  editCode() {
    this.mainService.updateCode(this.solution);
  }

}
