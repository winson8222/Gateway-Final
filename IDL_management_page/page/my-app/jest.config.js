module.exports = {
    roots: ["<rootDir>/tests"],
    transform: {
      "^.+\\.tsx?$": "babel-jest",
    },
    testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.tsx?$",
    moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json", "node"],
    collectCoverage: true,
    coverageReporters: ["json", "lcov", "text", "clover"],
    moduleNameMapper: {
      "\\.(css|less|sass|scss)$": "<rootDir>/__mocks__/styleMock.js",
      "\\.(gif|ttf|eot|svg|png)$": "<rootDir>/__mocks__/fileMock.js",
    },
    testEnvironment: 'jest-environment-jsdom',
    setupFilesAfterEnv: ['./jest.setup.js'],
  };
  