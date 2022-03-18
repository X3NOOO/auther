# Copyright © 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]
# 
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
# 
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.

bin_path = ./bin
bin_name = auther

user = $$(whoami)
user_install_path = $(HOME)/.local/bin
root_install_path = /usr/local/bin

release:
	go build -o $(bin_path)/$(bin_name) main.go

install:
	echo $(user_install_path);
	if test $(user) != "root"; \
        then \
            echo "You are not root. Installing to user path $(user_install_path)"; \
			cp $(bin_path)/$(bin_name) $(user_install_path); \
        else \
            echo "Installing to $(user_install_path)"; \
        	cp $(bin_path)/$(bin_name) $(root_install_path); \
        fi

uninstall:
	echo $(user_install_path);
	if test $(user) != "root"; \
        then \
            echo "You are not root. Uninstalling from user path $(user_install_path)"; \
			rm $(user_install_path)/$(bin_name); \
        else \
            echo "Uninstalling from $(user_install_path)"; \
        	rm $(root_install_path)/$(bin_name); \
        fi
