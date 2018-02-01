export class TestResult {
    testsPassed: boolean;
    buildfailed: boolean;
    builderr:   string[];
    failedtests: Test[];
}

export class Test {
    Expected: string;
    Got: string;
}
