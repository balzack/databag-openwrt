include $(TOPDIR)/rules.mk

PKG_NAME:=databag
PKG_VERSION:=1.1.26
PKG_RELEASE:=1

PKG_SOURCE:=v$(PKG_VERSION).tar.gz
PKG_SOURCE_URL:=https://github.com/balzack/databag-openwrt/archive/refs/tags
PKG_HASH:=cc42b1817cf30ef8f2bf6e7305f1a940d07a636106df7f92e46dcb3b7e49640b

PKG_MAINTAINER:=Roland Osborne <roland.osborne@gmail.com>
PKG_LICENSE:=Apache-2.0
PKG_LICENSE_FILES:=LICENSE

PKG_BUILD_DIR:=$(BUILD_DIR)/databag-openwrt-$(PKG_VERSION)
PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_BUILD_FLAGS:=no-mips16

GO_PKG:=databag

include $(INCLUDE_DIR)/package.mk
include ../../lang/golang/golang-package.mk

define Package/databag/install
	$(call GoPackage/Package/Install/Bin,$(1))
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(GO_PKG_BUILD_BIN_DIR)/databag $(1)/usr/bin/databag
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/databag.init $(1)/etc/init.d/databag
	$(INSTALL_DIR) $(1)/etc/databag
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/store_path.txt $(1)/etc/databag
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/service_port.txt $(1)/etc/databag
	$(INSTALL_DIR) $(1)/opt/databag
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/databag.db $(1)/opt/databag/
	$(CP) -r $(PKG_BUILD_DIR)/web $(1)/opt/databag/
endef

define Package/databag
  SECTION:=net
  CATEGORY:=Network
  SUBMENU:=Instant Messaging
  TITLE:=Federated Messaging Service
  URL:=https://github.com/balzack/databag
  DEPENDS:=$(GO_ARCH_DEPENDS)
endef

define Package/databagdescription
  Databag is a federated messaging service with browser and mobile clients.
endef

$(eval $(call GoBinPackage,databag))
$(eval $(call BuildPackage,databag))

