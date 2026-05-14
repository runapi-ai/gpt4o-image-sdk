# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    class Client
      attr_reader :generations

      def initialize(api_key: nil, **options)
        @api_key = Core::Auth.resolve_api_key(api_key)

        client_options = Core::ClientOptions.new(api_key: @api_key, **options)
        http = client_options.http_client || Core::HttpClient.new(client_options)
        @generations = Resources::Generations.new(http)
      end
    end
  end
end
