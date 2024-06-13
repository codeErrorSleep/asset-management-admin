/**********************************
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/05 21:24:46
 * @Email: zclzone@outlook.com
 *  | https://isme.top
 **********************************/

import { createPageLoadingGuard } from './page-loading-guard'
import { createPageTitleGuard } from './page-title-guard'
import { createPermissionGuard } from './permission-guard'
import { createTabGuard } from './tab-guard'

export function setupRouterGuards(router) {
  createPageLoadingGuard(router)
  createPermissionGuard(router)
  createPageTitleGuard(router)
  createTabGuard(router)
}
