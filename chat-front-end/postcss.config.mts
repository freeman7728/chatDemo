/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-24 11:06:43
 * @LastEditTime: 2024-09-24 11:06:50
 * @LastEditors: freeman7728
 */
export const defaultHtmlFontSize = 37.5
export default {
  plugins: {
    autoprefixer: {
      overrideBrowserslist: ['Android >= 4.0', 'iOS >= 7'],
    },
    'postcss-pxtorem': {
      // 根节点的 fontSize 值
      rootValue: defaultHtmlFontSize,
      propList: ['*'],
      selectorBlackList: [':root'],
    },
  },
}