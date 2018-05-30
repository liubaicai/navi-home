# frozen_string_literal: true

class Users::RegistrationsController < Devise::RegistrationsController
  # before_action :configure_sign_up_params, only: [:create]
  # before_action :configure_account_update_params, only: [:update]

  # GET /resource/sign_up
  # def new
  #   super
  # end

  # POST /resource
  def create
    build_resource(sign_up_params)

    resource.save
    yield resource if block_given?
    if resource.persisted?
      catalog1 = Catalog.create(title: '综合', sort_by:1, user_id:resource.id);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog1.id, sort_by:1);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog1.id, sort_by:2);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog1.id, sort_by:3);
      catalog2 = Catalog.create(title: '社区', sort_by:2, user_id:resource.id);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog2.id, sort_by:1);
      catalog3 = Catalog.create(title: '影视', sort_by:3, user_id:resource.id);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog3.id, sort_by:1);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog3.id, sort_by:2);
      catalog4 = Catalog.create(title: '工具', sort_by:4, user_id:resource.id);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:1);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:2);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:3);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:4);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:5);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog4.id, sort_by:6);
      catalog5 = Catalog.create(title: '购物', sort_by:5, user_id:resource.id);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog5.id, sort_by:1);
      Link.create(title: '微博', url: 'http://weibo.com/?is_search=1',icon: 'https://ss2.bdstatic.com/lfoZeXSm1A5BphGlnYG/icon/54.png?1',
                  user_id:resource.id, catalog_id:catalog5.id, sort_by:2);
      if resource.active_for_authentication?
        set_flash_message! :notice, :signed_up
        sign_up(resource_name, resource)
        respond_with resource, location: after_sign_up_path_for(resource)
      else
        set_flash_message! :notice, :"signed_up_but_#{resource.inactive_message}"
        expire_data_after_sign_in!
        respond_with resource, location: after_inactive_sign_up_path_for(resource)
      end
    else
      clean_up_passwords resource
      set_minimum_password_length
      respond_with resource
    end
  end

  # GET /resource/edit
  # def edit
  #   super
  # end

  # PUT /resource
  # def update
  #   super
  # end

  # DELETE /resource
  # def destroy
  #   super
  # end

  # GET /resource/cancel
  # Forces the session data which is usually expired after sign
  # in to be expired now. This is useful if the user wants to
  # cancel oauth signing in/up in the middle of the process,
  # removing all OAuth session data.
  # def cancel
  #   super
  # end

  # protected

  # If you have extra params to permit, append them to the sanitizer.
  # def configure_sign_up_params
  #   devise_parameter_sanitizer.permit(:sign_up, keys: [:attribute])
  # end

  # If you have extra params to permit, append them to the sanitizer.
  # def configure_account_update_params
  #   devise_parameter_sanitizer.permit(:account_update, keys: [:attribute])
  # end

  # The path used after sign up.
  # def after_sign_up_path_for(resource)
  #   super(resource)
  # end

  # The path used after sign up for inactive accounts.
  # def after_inactive_sign_up_path_for(resource)
  #   super(resource)
  # end
end
