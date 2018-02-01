import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Subject } from 'rxjs/Subject';
import { environment as env } from '../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { HttpResponse } from 'selenium-webdriver/http';

import { Test, TestResult } from '../models/testresult';

@Injectable()
export class MainService {
  private CodeSolutionSource = new Subject<string>();
  solution: string;

  CodeSolution$ = this.CodeSolutionSource.asObservable();

  constructor(private http: HttpClient) { }

  updateCode (code: string) {
    this.CodeSolutionSource.next(code);
    this.CodeSolution$.subscribe(data => this.solution = data);
  }

  runTests(challenge, code: string): Observable<HttpResponse> {
    console.log(code);
    return this.http
      .post( env.apiURL + `runtests/${challenge}`, code, { observe: 'response' });
  }

  fmt(code: string) {
    return this.http
      .post( env.apiURL + 'fmt', code, { observe: 'response' })
      .subscribe(data => console.log(data), err => console.log(err));
  }

}
