class HomeController < ApplicationController
  def index
    @catalogs = current_user.catalogs.as_json(:include => 'links').to_json
  end

  def set_catalogs

  end

end
