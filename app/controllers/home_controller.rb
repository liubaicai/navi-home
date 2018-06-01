require 'open-uri'
require 'net/http'

class HomeController < ApplicationController
  def index
    if current_user
      @catalogs = current_user.catalogs.as_json(:include => 'links').to_json
    end
  end

  def test_icon_url
    uri = URI(params["url"])
    Net::HTTP.start(uri.host, uri.port, use_ssl: uri.scheme == 'https') do |http|
      request = Net::HTTP::Head.new uri
      http.request request do |response|
        if response.code.to_i>299
          render plain: '/ie.png'
        else
          render plain: params["url"]
        end
      end
    end
  end

  def set_catalogs
    if current_user
      catalogs = JSON.parse(params["data"])

      catalogs_id = Array.new
      catalogs_del = Array.new
      catalogs.each do |item|
        catalogs_id.push item['id']
      end
      current_user.catalogs.each do |catalog|
        unless catalogs_id.include? catalog.id
          catalogs_del.push catalog
        end
      end

      catalogs.each_with_index do |item,index|
        if item['id'] >= 0
          catalog = Catalog.where(:id => item['id'],:user_id => current_user.id).first
          catalog.title = item['title']
          catalog.sort_by = index+1
          catalog.save

          links_id = Array.new
          links_del = Array.new
          item['links'].each do |link|
            links_id.push link['id']
          end
          catalog.links.each do |link|
            unless links_id.include? link.id
              links_del.push link
            end
          end

          item['links'].each_with_index do |link2,index2|
            if link2['id'] >= 0
              link = Link.where(:id => link2['id'],:user_id => current_user.id).first
              link.title = link2['title']
              link.url = link2['url']
              link.icon = link2['icon']
              link.sort_by = index2+1
              link.save
            else
              Link.create(title: link2['title'], url: link2['url'],icon: link2['icon'],
                          user_id:current_user.id, catalog_id:catalog.id, sort_by:(index2+1));
            end
          end

          links_del.each do |link|
            link.destroy
          end

        else
          catalog = Catalog.create(title: item['title'], sort_by:(index+1), user_id:current_user.id);

          item['links'].each_with_index do |link,index2|
            Link.create(title: link['title'], url: link['url'],icon: link['icon'],
                        user_id:current_user.id, catalog_id:catalog.id, sort_by:(index2+1));
          end

        end
      end

      catalogs_del.each do |catalog|
        catalog.destroy
      end

      render plain: 'ok'
    else
      render plain: 'not ok'
    end
  end

end
