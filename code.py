import os

# 読み込みたいファイルの拡張子を指定
target_extensions = ['.jsx', '.js', '.css', '.go', '.yaml'] 
# 除外したいフォルダ
ignore_dirs = ['.git', '__pycache__', '.idea', 'venv', 'node_modules']

def print_all_codes():
    print("=== PROJECT CODE START ===")
    
    # 現在のフォルダ以下の全ファイルを探索
    for root, dirs, files in os.walk("."):
        # 除外フォルダを無視する設定
        dirs[:] = [d for d in dirs if d not in ignore_dirs]
        
        for file in files:
            # 拡張子をチェック
            if any(file.endswith(ext) for ext in target_extensions):
                file_path = os.path.join(root, file)
                
                print(f"\n\n--- FILE: {file_path} ---\n")
                
                try:
                    with open(file_path, 'r', encoding='utf-8') as f:
                        print(f.read())
                except Exception as e:
                    print(f"Error reading file: {e}")

    print("\n=== PROJECT CODE END ===")

if __name__ == "__main__":
    print_all_codes()

# aaaaaaaa
# bbbbbbbb
# cccccccc
# dddddddd