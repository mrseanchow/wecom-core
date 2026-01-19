#!/usr/bin/env python3
"""Web to Markdown Agent

功能：
1. 从 URL 抓取网页内容
2. 转换成 Markdown 格式
3. 保存到本地文件

python3 -m venv venv

source venv/bin/activate 
依赖安装：
pip install -r requirements.txt 

运行：
python3 web_to_markdown_agent.py
"""

import os
import requests
from bs4 import BeautifulSoup
import html2text
import csv
import time
from datetime import datetime
from dotenv import load_dotenv


def fetch_webpage(url: str) -> str:
    """从给定的URL抓取网页内容"""
    try:
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        }
        print(f"\n正在抓取网页: {url}")
        response = requests.get(url, headers=headers, timeout=30)
        response.raise_for_status()
        response.encoding = response.apparent_encoding
        print("✓ 网页抓取成功")
        return response.text
    except Exception as e:
        error_msg = f"Error fetching webpage: {str(e)}"
        print(f"✗ {error_msg}")
        return error_msg


def convert_html_to_markdown(html_content: str) -> str:
    """将HTML内容转换为Markdown格式"""
    try:
        print("\n正在转换HTML到Markdown...")
        # 使用 BeautifulSoup 解析 HTML
        soup = BeautifulSoup(html_content, 'html.parser')

        # 查找 id=js-ep-doc-area 的 div
        target_div = soup.find('div', id='js-ep-doc-cnt')

        if not target_div:
            error_msg = "未找到 id='js-ep-doc-area' 的 div 元素"
            print(f"✗ {error_msg}")
            return error_msg

        print("✓ 找到目标内容区域")

        # 移除目标 div 中的 script 和 style 标签
        for script in target_div(['script', 'style']):
            script.decompose()

        # 转换为 Markdown
        h = html2text.HTML2Text()
        h.ignore_links = False
        h.ignore_images = False
        h.ignore_emphasis = False
        h.body_width = 0  # 不自动换行

        markdown_content = h.handle(str(target_div))
        print("✓ 转换成功")
        return markdown_content
    except Exception as e:
        error_msg = f"Error converting to markdown: {str(e)}"
        print(f"✗ {error_msg}")
        return error_msg


def parse_ep_doc_items(html_content: str, base_url: str = "") -> list:
    """解析HTML内容中div id=js-ep-sidebar内带有class="ep-doc-item"的a标签

    Args:
        html_content: HTML内容
        base_url: 基础URL，用于拼接相对路径

    Returns:
        包含字典的列表，每个字典包含 'title' 和 'url' 键
    """
    try:
        print("\n正在解析 ep-doc-item 链接...")
        soup = BeautifulSoup(html_content, 'html.parser')

        # 查找 id=js-ep-sidebar 的 div
        sidebar_div = soup.find('div', id='js-ep-sidebar')

        if not sidebar_div:
            print("✗ 未找到 id='js-ep-sidebar' 的 div 元素")
            return []

        print("✓ 找到侧边栏区域")

        # 在侧边栏内查找所有带有 class="ep-doc-item" 的 a 标签
        doc_items = sidebar_div.find_all('a', class_='ep-doc-item')

        results = []
        for item in doc_items:
            title = item.get_text(strip=True)
            href = item.get('href', '')

            # 如果是相对路径且提供了base_url，则拼接完整URL
            if href and base_url and not href.startswith(('http://', 'https://')):
                if not base_url.endswith('/'):
                    base_url += '/'
                if href.startswith('/'):
                    href = href[1:]
                full_url = base_url + href
            else:
                full_url = href

            if title and href:
                results.append({
                    'title': title,
                    'url': full_url
                })

        print(f"✓ 找到 {len(results)} 个 ep-doc-item 链接")
        return results
    except Exception as e:
        print(f"✗ 解析错误: {str(e)}")
        return []


def save_to_csv(data: list, csv_filename: str = "企业微信接口文档.csv") -> str:
    """将解析的链接数据保存到CSV文件

    Args:
        data: 包含字典的列表，每个字典包含 'title' 和 'url' 键
        csv_filename: CSV文件名

    Returns:
        保存的文件路径
    """
    try:
        # 获取当前脚本所在目录
        base_dir = os.path.dirname(os.path.abspath(__file__))
        csv_filepath = os.path.join(base_dir, csv_filename)

        print(f"\n正在保存到CSV文件: {csv_filepath}")

        # 写入CSV文件
        with open(csv_filepath, 'w', encoding='utf-8-sig', newline='') as f:
            writer = csv.DictWriter(f, fieldnames=['title', 'url'])
            writer.writeheader()
            writer.writerows(data)

        print(f"✓ 成功保存 {len(data)} 条记录到: {csv_filepath}")
        return csv_filepath
    except Exception as e:
        error_msg = f"Error saving CSV: {str(e)}"
        print(f"✗ {error_msg}")
        return error_msg


def save_to_file(content: str, filename: str = "") -> str:
    """保存内容到本地文件"""
    try:
        # 创建保存目录
        save_dir = os.path.join(
            os.path.dirname(os.path.abspath(__file__)),
            'markdown'
        )
        os.makedirs(save_dir, exist_ok=True)

        # 生成文件名
        if not filename:
            timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
            filename = f'webpage_{timestamp}.md'
        elif not filename.endswith('.md'):
            filename += '.md'

        # 保存文件
        filepath = os.path.join(save_dir, filename)

        print(f"\n正在保存文件: {filepath}")
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)

        success_msg = f"✓ 成功保存到: {filepath}"
        print(success_msg)
        return filepath
    except Exception as e:
        error_msg = f"Error saving file: {str(e)}"
        print(f"✗ {error_msg}")
        return error_msg


def process_url(url: str, filename: str = "") -> str:
    """处理URL的完整流程：抓取 -> 转换 -> 保存"""
    print("=" * 80)
    print("开始处理网页...")
    print("=" * 80)

    # 1. 抓取网页
    html_content = fetch_webpage(url)
    if html_content.startswith("Error"):
        return html_content

    # 2. 转换为Markdown
    markdown_content = convert_html_to_markdown(html_content)
    if markdown_content.startswith("Error"):
        return markdown_content

    # 3. 保存文件
    filepath = save_to_file(markdown_content, filename)

    print("\n" + "=" * 80)
    print("处理完成！")
    print("=" * 80)

    return filepath


def read_csv_and_process(csv_filename: str = "企业微信接口文档.csv") -> None:
    """读取CSV文件中的URL列表，并批量抓取转换为Markdown

    Args:
        csv_filename: CSV文件名
    """
    try:
        # 获取当前脚本所在目录
        base_dir = os.path.dirname(os.path.abspath(__file__))
        csv_filepath = os.path.join(base_dir, csv_filename)

        # 检查CSV文件是否存在
        if not os.path.exists(csv_filepath):
            print(f"✗ CSV文件不存在: {csv_filepath}")
            return

        print("=" * 80)
        print(f"开始读取CSV文件: {csv_filepath}")
        print("=" * 80)

        # 读取CSV文件
        with open(csv_filepath, 'r', encoding='utf-8-sig') as f:
            reader = csv.DictReader(f)
            rows = list(reader)

        if not rows:
            print("✗ CSV文件为空")
            return

        print(f"\n✓ 找到 {len(rows)} 条记录")
        print("\n开始批量处理...\n")

        success_count = 0
        fail_count = 0
        skip_count = 0

        # 获取保存目录
        save_dir = os.path.join(base_dir, 'data', 'markdown_output')
        os.makedirs(save_dir, exist_ok=True)

        # 逐条处理
        for index, row in enumerate(rows, 1):
            title = row.get('title', '').strip()
            url = row.get('url', '').strip()

            if not url:
                print(f"\n[{index}/{len(rows)}] ✗ 跳过：URL为空")
                skip_count += 1
                continue

            # 生成文件名
            filename = title if title else f"webpage_{index}"
            if not filename.endswith('.md'):
                filename += '.md'

            # 检查文件是否已存在
            filepath = os.path.join(save_dir, filename)
            if os.path.exists(filepath):
                print(f"\n[{index}/{len(rows)}] ⊙ 跳过: {title} (文件已存在)")
                skip_count += 1
                continue

            print(f"\n{'=' * 80}")
            print(f"[{index}/{len(rows)}] 处理: {title}")
            print(f"URL: {url}")
            print('=' * 80)

            try:
                # 处理URL
                result = process_url(url, title if title else f"webpage_{index}")

                if result and not result.startswith("Error"):
                    success_count += 1
                    print(f"✓ [{index}/{len(rows)}] 成功")
                else:
                    fail_count += 1
                    print(f"✗ [{index}/{len(rows)}] 失败")

            except Exception as e:
                fail_count += 1
                print(f"✗ [{index}/{len(rows)}] 错误: {str(e)}")

            # 速率控制：每次请求后等待1秒（避免被服务器封禁）
            if index < len(rows):  # 最后一条记录不需要等待
                print("等待1秒...")
                time.sleep(1)

        # 打印统计信息
        print("\n" + "=" * 80)
        print("批量处理完成！")
        print("=" * 80)
        print(f"总计: {len(rows)} 条")
        print(f"成功: {success_count} 条")
        print(f"失败: {fail_count} 条")
        print(f"跳过: {skip_count} 条（已存在或URL为空）")
        print("=" * 80)

    except Exception as e:
        print(f"✗ 读取CSV文件错误: {str(e)}")


def main():
    # 加载环境变量
    base_dir = os.path.dirname(os.path.abspath(__file__))
    dotenv_path = os.path.join(base_dir, '.env')
    load_dotenv(dotenv_path)

    # 主菜单
    print("=" * 80)
    print("欢迎使用 Web to Markdown 工具")
    print("=" * 80)
    print("\n请选择模式：")
    print("1. 交互模式 - 手动输入URL")
    print("2. 批量模式 - 从CSV文件读取URL列表")
    print()

    mode = input("请选择模式 (1 或 2): ").strip()

    if mode == '2':
        # 批量模式
        csv_file = input("\n请输入CSV文件名 (直接回车使用默认 '企业微信接口文档.csv'): ").strip()
        if not csv_file:
            csv_file = "企业微信接口文档.csv"
        read_csv_and_process(csv_file)
    else:
        # 交互模式
        print("\n" + "=" * 80)
        print("交互模式")
        print("=" * 80)
        print("\n这个工具可以帮你：")
        print("1. 抓取任何网页的内容")
        print("2. 自动转换为 Markdown 格式")
        print("3. 保存到本地文件")
        print()

        while True:
            url = input("\n请输入要抓取的网页URL (输入 'quit' 或 'q' 退出): ").strip()
            if url.lower() in ['quit', 'q']:
                print("\n再见！")
                break

            if not url:
                print("✗ URL 不能为空")
                continue

            if not url.startswith(('http://', 'https://')):
                print("✗ URL 必须以 http:// 或 https:// 开头")
                continue

            filename = input("请输入保存的文件名 (直接回车使用时间戳命名): ").strip()

            try:
                result = process_url(url, filename)
                print(f"\n文件已保存: {result}")
            except Exception as e:
                print(f"\n✗ 错误: {str(e)}")


if __name__ == '__main__':
    main()
