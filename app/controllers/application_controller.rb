class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  before_filter :mailer_set_url_options
  before_action :configure_permitted_parameters, if: :devise_controller?

  def configure_permitted_parameters
    added_attrs = [:username, :email, :password, :password_confirmation, :remember_me]
    devise_parameter_sanitizer.permit :sign_up, keys: added_attrs
    devise_parameter_sanitizer.permit :account_update, keys: added_attrs
  end

  def mailer_set_url_options
    ActionMailer::Base.default_url_options = {:host => request.host_with_port}
  end
end
