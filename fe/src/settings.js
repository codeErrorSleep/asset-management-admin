/**********************************
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/13 20:54:36
 * @Email: zclzone@outlook.com
 *  | https://isme.top
 **********************************/

export const defaultLayout = 'normal'

export const defaultPrimaryColor = '#316C72'

// 控制 LayoutSetting 组件是否可见
export const layoutSettingVisible = true

export const naiveThemeOverrides = {
  common: {
    primaryColor: '#316C72FF',
    primaryColorHover: '#316C72E3',
    primaryColorPressed: '#2B4C59FF',
    primaryColorSuppl: '#316C72E3',
  },
}

export const basePermissions = [
  {
    code: 'base-project',
    name: '配网项目管理',
    type: 'MENU',
    icon: 'i-fe:external-link',
    order: 0,
    enable: true,
    show: true,
    children: [
      {
        code: 'tower-detail',
        name: '杆塔管理',
        type: 'MENU',
        path: '/project/tower',
        icon: 'i-me:docs',
        order: 1,
        enable: true,
        show: true,
        component: "/src/views/project/tower/index.vue",
      },
      {
        code: 'subitem-detail',
        name: '子项管理',
        type: 'MENU',
        path: '/base/components',
        icon: 'i-me:docs',
        order: 1,
        enable: true,
        show: true,
        component: "/src/views/base/unocss-icon.vue",
      },
      {
        code: 'project-package',
        name: '项目包管理',
        type: 'MENU',
        path: '/base/keep-alive',
        icon: 'i-me:docs',
        order: 1,
        enable: true,
        show: true,
        component: "/src/views/base/unocss-icon.vue",
      }
    ]
  },
  {
    code: 'equipment-management',
    name: '设备管理',
    type: 'MENU',
    icon: 'i-me:docs',
    order: 0,
    enable: true,
    show: true,
    children: [
      {
        code: 'equipment-detail',
        name: '设备管理',
        type: 'MENU',
        path: '/base/components',
        icon: 'i-fe:tool',
        order: 1,
        enable: true,
        show: true,
        component: "/src/views/equipment/detail.vue",
      },
      {
        code: 'equipment-class',
        name: '设备分类',
        type: 'MENU',
        path: '/base/components',
        icon: 'i-me:docs',
        order: 2,
        enable: true,
        show: true,
        component: "/src/views/equipment/class.vue",
      }
    ]
  },
  {
    "id": 9,
    "name": "基础功能",
    "code": "Base",
    "type": "MENU",
    "parentId": null,
    "path": "",
    "redirect": null,
    "icon": "i-fe:grid",
    "component": null,
    "layout": "",
    "keepAlive": null,
    "method": null,
    "description": null,
    "show": true,
    "enable": true,
    "order": 0,
    "children": [
      {
        "id": 14,
        "name": "图标 Icon",
        "code": "Icon",
        "type": "MENU",
        "parentId": 9,
        "path": "/base/icon",
        "redirect": null,
        "icon": "i-fe:feather",
        "component": "/src/views/base/unocss-icon.vue",
        "layout": "",
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 0
      },
      {
        "id": 10,
        "name": "基础组件",
        "code": "BaseComponents",
        "type": "MENU",
        "parentId": 9,
        "path": "/base/components",
        "redirect": null,
        "icon": "i-me:awesome",
        "component": "/src/views/base/index.vue",
        "layout": null,
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 1
      },
      {
        "id": 11,
        "name": "Unocss",
        "code": "Unocss",
        "type": "MENU",
        "parentId": 9,
        "path": "/base/unocss",
        "redirect": null,
        "icon": "i-me:awesome",
        "component": "/src/views/base/unocss.vue",
        "layout": null,
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 2
      },
      {
        "id": 12,
        "name": "KeepAlive",
        "code": "KeepAlive",
        "type": "MENU",
        "parentId": 9,
        "path": "/base/keep-alive",
        "redirect": null,
        "icon": "i-me:awesome",
        "component": "/src/views/base/keep-alive.vue",
        "layout": null,
        "keepAlive": true,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 3
      },
      {
        "id": 15,
        "name": "MeModal",
        "code": "TestModal",
        "type": "MENU",
        "parentId": 9,
        "path": "/testModal",
        "redirect": null,
        "icon": "i-me:dialog",
        "component": "/src/views/base/test-modal.vue",
        "layout": null,
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 5
      }
    ]
  },
  {
    "id": 6,
    "name": "业务示例",
    "code": "Demo",
    "type": "MENU",
    "parentId": null,
    "path": null,
    "redirect": null,
    "icon": "i-fe:grid",
    "component": null,
    "layout": null,
    "keepAlive": null,
    "method": null,
    "description": null,
    "show": true,
    "enable": true,
    "order": 1,
    "children": [
      {
        "id": 7,
        "name": "图片上传",
        "code": "ImgUpload",
        "type": "MENU",
        "parentId": 6,
        "path": "/demo/upload",
        "redirect": null,
        "icon": "i-fe:image",
        "component": "/src/views/demo/upload/index.vue",
        "layout": "simple",
        "keepAlive": true,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 2
      }
    ]
  },
  {
    "id": 2,
    "name": "系统管理",
    "code": "SysMgt",
    "type": "MENU",
    "parentId": null,
    "path": null,
    "redirect": null,
    "icon": "i-fe:grid",
    "component": null,
    "layout": null,
    "keepAlive": null,
    "method": null,
    "description": null,
    "show": true,
    "enable": true,
    "order": 2,
    "children": [
      {
        "id": 1,
        "name": "资源管理",
        "code": "Resource_Mgt",
        "type": "MENU",
        "parentId": 2,
        "path": "/pms/resource",
        "redirect": null,
        "icon": "i-fe:list",
        "component": "/src/views/pms/resource/index.vue",
        "layout": null,
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 1
      },
      {
        "id": 3,
        "name": "角色管理",
        "code": "RoleMgt",
        "type": "MENU",
        "parentId": 2,
        "path": "/pms/role",
        "redirect": null,
        "icon": "i-fe:user-check",
        "component": "/src/views/pms/role/index.vue",
        "layout": null,
        "keepAlive": null,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 2,
        "children": [
          {
            "id": 5,
            "name": "分配用户",
            "code": "RoleUser",
            "type": "MENU",
            "parentId": 3,
            "path": "/pms/role/user/:roleId",
            "redirect": null,
            "icon": "i-fe:user-plus",
            "component": "/src/views/pms/role/role-user.vue",
            "layout": "full",
            "keepAlive": null,
            "method": null,
            "description": null,
            "show": false,
            "enable": true,
            "order": 1
          }
        ]
      },
      {
        "id": 4,
        "name": "用户管理",
        "code": "UserMgt",
        "type": "MENU",
        "parentId": 2,
        "path": "/pms/user",
        "redirect": null,
        "icon": "i-fe:user",
        "component": "/src/views/pms/user/index.vue",
        "layout": null,
        "keepAlive": true,
        "method": null,
        "description": null,
        "show": true,
        "enable": true,
        "order": 3
      }
    ]
  },
  {
    "id": 8,
    "name": "个人资料",
    "code": "UserProfile",
    "type": "MENU",
    "parentId": null,
    "path": "/profile",
    "redirect": null,
    "icon": "i-fe:user",
    "component": "/src/views/profile/index.vue",
    "layout": null,
    "keepAlive": null,
    "method": null,
    "description": null,
    "show": false,
    "enable": true,
    "order": 99
  },
  {
    code: 'ExternalLink',
    name: '外链(可内嵌打开)',
    type: 'MENU',
    icon: 'i-fe:external-link',
    order: 98,
    enable: true,
    show: true,
    children: [
      {
        code: 'ShowDocs',
        name: '项目文档',
        type: 'MENU',
        path: 'https://docs.isme.top/web/#/624306705/188522224',
        icon: 'i-me:docs',
        order: 1,
        enable: true,
        show: true,
      },
      {
        code: 'ApiFoxDocs',
        name: '接口文档',
        type: 'MENU',
        path: 'https://apifox.com/apidoc/shared-ff4a4d32-c0d1-4caf-b0ee-6abc130f734a',
        icon: 'i-me:apifox',
        order: 2,
        enable: true,
        show: true,
      },
      {
        code: 'NaiveUI',
        name: 'Naive UI',
        type: 'MENU',
        path: 'https://www.naiveui.com/zh-CN/os-theme',
        icon: 'i-me:naiveui',
        order: 3,
        enable: true,
        show: true,
      },
      {
        code: 'MyBlog',
        name: '博客-掘金',
        type: 'MENU',
        path: 'https://juejin.cn/user/1961184475483255/posts',
        icon: 'i-simple-icons:juejin',
        order: 4,
        enable: true,
        show: true,
      },
    ],
  },
]
