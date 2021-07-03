module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/essential',
    '@vue/standard'
  ],
  parserOptions: {
    parser: 'babel-eslint'
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    "indent": ["off", 2],
    // "space-before-function-paren":  ["error" ,{
    //   anonymous: 'always',
    //   named: 'always',
    //   asyncArrow: 'always'
    // }],
    "key-spacing": 0,
    "spack-before-blocks": 0,
    "space-before-function-paren": 0,
    "arrow-spacing": 0,
    "comma-spacing": 0
  },
  globals: {
    'config': true
  }
}
