module.exports = {
  extends: ['plugin:@typescript-eslint/recommended', 'prettier'],
  rules: {
    quotes: [
      'error',
      'single',
      { avoidEscape: true, allowTemplateLiterals: false },
    ],
  },
};
